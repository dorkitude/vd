package search

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"unicode"
)

// Result represents a search result
type Result struct {
	Project     string `json:"project"`
	File        string `json:"file"`
	Path        string `json:"path"`
	Line        int    `json:"line"`
	Column      int    `json:"column"`
	Match       string `json:"match"`
	Context     string `json:"context"`
	Score       int    `json:"score"`
	BeforeLines []string `json:"before_lines,omitempty"`
	AfterLines  []string `json:"after_lines,omitempty"`
}

// SearchOptions configures the search behavior
type SearchOptions struct {
	Project       string
	Query         string
	CaseSensitive bool
	ContextLines  int
	MaxResults    int
	FuzzyMatch    bool
	OutputFormat  string // "text" or "json"
}

// Engine handles document searching
type Engine struct {
	contentDir string
	options    SearchOptions
}

// NewEngine creates a new search engine
func NewEngine(contentDir string) *Engine {
	return &Engine{
		contentDir: contentDir,
	}
}

// Search performs a search across documents
func (e *Engine) Search(opts SearchOptions) ([]Result, error) {
	var results []Result
	
	// Normalize query for searching
	query := opts.Query
	if !opts.CaseSensitive {
		query = strings.ToLower(query)
	}
	
	// Determine which projects to search
	projects, err := e.getProjects(opts.Project)
	if err != nil {
		return nil, err
	}
	
	// Search each project
	for _, project := range projects {
		projectResults, err := e.searchProject(project, query, opts)
		if err != nil {
			continue // Skip failed projects
		}
		results = append(results, projectResults...)
		
		// Limit total results
		if opts.MaxResults > 0 && len(results) >= opts.MaxResults {
			results = results[:opts.MaxResults]
			break
		}
	}
	
	// Sort by relevance score
	e.sortByScore(results)
	
	return results, nil
}

// getProjects returns the list of projects to search
func (e *Engine) getProjects(projectFilter string) ([]string, error) {
	var projects []string
	
	entries, err := ioutil.ReadDir(e.contentDir)
	if err != nil {
		return nil, err
	}
	
	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}
		
		// Check if this is a valid project with vd.json
		metadataPath := filepath.Join(e.contentDir, entry.Name(), "vd.json")
		if _, err := os.Stat(metadataPath); err != nil {
			continue
		}
		
		// Apply project filter if specified
		if projectFilter != "" {
			if !strings.Contains(strings.ToLower(entry.Name()), strings.ToLower(projectFilter)) {
				continue
			}
		}
		
		projects = append(projects, entry.Name())
	}
	
	return projects, nil
}

// searchProject searches within a single project
func (e *Engine) searchProject(project, query string, opts SearchOptions) ([]Result, error) {
	var results []Result
	projectPath := filepath.Join(e.contentDir, project)
	
	// Walk through all markdown files
	err := filepath.Walk(projectPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil // Skip files with errors
		}
		
		// Only search markdown files
		if !strings.HasSuffix(path, ".md") {
			return nil
		}
		
		// Search this file
		fileResults, err := e.searchFile(project, path, query, opts)
		if err != nil {
			return nil // Skip files with errors
		}
		
		results = append(results, fileResults...)
		return nil
	})
	
	if err != nil {
		return nil, err
	}
	
	return results, nil
}

// searchFile searches within a single file
func (e *Engine) searchFile(project, filePath, query string, opts SearchOptions) ([]Result, error) {
	var results []Result
	
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	
	scanner := bufio.NewScanner(file)
	lineNum := 0
	var lines []string
	
	// Read all lines for context
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	
	// Search through lines
	for i, line := range lines {
		lineNum = i + 1
		searchLine := line
		if !opts.CaseSensitive {
			searchLine = strings.ToLower(line)
		}
		
		// Check for match
		var matchIndex int
		var matched bool
		
		if opts.FuzzyMatch {
			matched, matchIndex = e.fuzzyMatch(searchLine, query)
		} else {
			matchIndex = strings.Index(searchLine, query)
			matched = matchIndex >= 0
		}
		
		if matched {
			// Calculate relative path
			relPath, _ := filepath.Rel(filepath.Join(e.contentDir, project), filePath)
			
			// Get context
			beforeLines := e.getContextLines(lines, i, opts.ContextLines, true)
			afterLines := e.getContextLines(lines, i, opts.ContextLines, false)
			
			// Create context string
			contextStart := maxInt(0, matchIndex-50)
			contextEnd := minInt(len(line), matchIndex+len(query)+50)
			context := line[contextStart:contextEnd]
			if contextStart > 0 {
				context = "..." + context
			}
			if contextEnd < len(line) {
				context = context + "..."
			}
			
			result := Result{
				Project:     project,
				File:        filepath.Base(filePath),
				Path:        relPath,
				Line:        lineNum,
				Column:      matchIndex + 1,
				Match:       line,
				Context:     context,
				Score:       e.calculateScore(query, line, matchIndex),
				BeforeLines: beforeLines,
				AfterLines:  afterLines,
			}
			
			results = append(results, result)
		}
	}
	
	return results, nil
}

// fuzzyMatch performs fuzzy string matching
func (e *Engine) fuzzyMatch(text, pattern string) (bool, int) {
	// Simple fuzzy match: all pattern characters must appear in order
	patternIdx := 0
	firstMatch := -1
	
	for i, ch := range text {
		if patternIdx >= len(pattern) {
			break
		}
		
		if ch == rune(pattern[patternIdx]) {
			if firstMatch == -1 {
				firstMatch = i
			}
			patternIdx++
		}
	}
	
	return patternIdx == len(pattern), firstMatch
}

// calculateScore calculates relevance score for a match
func (e *Engine) calculateScore(query, line string, position int) int {
	score := 100
	
	// Exact match gets highest score
	if strings.Contains(strings.ToLower(line), strings.ToLower(query)) {
		score += 50
	}
	
	// Match at beginning of line scores higher
	if position == 0 {
		score += 30
	} else if position < 10 {
		score += 20
	}
	
	// Match at word boundary scores higher
	if position > 0 && unicode.IsSpace(rune(line[position-1])) {
		score += 25
	}
	
	// Matches in headers score higher
	if strings.HasPrefix(line, "#") {
		score += 40
	}
	
	return score
}

// getContextLines retrieves context lines before or after a match
func (e *Engine) getContextLines(lines []string, index, count int, before bool) []string {
	var context []string
	
	if before {
		start := maxInt(0, index-count)
		for i := start; i < index; i++ {
			context = append(context, lines[i])
		}
	} else {
		end := minInt(len(lines), index+count+1)
		for i := index + 1; i < end; i++ {
			context = append(context, lines[i])
		}
	}
	
	return context
}

// sortByScore sorts results by relevance score
func (e *Engine) sortByScore(results []Result) {
	// Simple bubble sort for now
	n := len(results)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if results[j].Score < results[j+1].Score {
				results[j], results[j+1] = results[j+1], results[j]
			}
		}
	}
}

// FormatResults formats search results based on output format
func FormatResults(results []Result, format string) string {
	if format == "json" {
		data, _ := json.MarshalIndent(results, "", "  ")
		return string(data)
	}
	
	// Text format
	var output strings.Builder
	
	if len(results) == 0 {
		output.WriteString("No results found.\n")
		return output.String()
	}
	
	output.WriteString(fmt.Sprintf("Found %d results:\n\n", len(results)))
	
	currentProject := ""
	for _, result := range results {
		// Group by project
		if result.Project != currentProject {
			currentProject = result.Project
			output.WriteString(fmt.Sprintf("\n📁 %s\n", result.Project))
			output.WriteString(strings.Repeat("-", 40) + "\n")
		}
		
		// Format result
		output.WriteString(fmt.Sprintf("  📄 %s:%d:%d\n", result.Path, result.Line, result.Column))
		output.WriteString(fmt.Sprintf("     %s\n", result.Context))
		output.WriteString("\n")
	}
	
	return output.String()
}

// Helper functions
func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}