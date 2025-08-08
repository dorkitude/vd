package search

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
)

// OutputMode defines how results are formatted
type OutputMode string

const (
	OutputFiles   OutputMode = "files_with_matches"
	OutputContent OutputMode = "content"
	OutputCount   OutputMode = "count"
)

// GrepOptions configures grep-like search behavior
type GrepOptions struct {
	Pattern      string     // Search pattern (regex or literal)
	Path         string     // Path to search in (project or specific dir)
	OutputMode   OutputMode  // How to format results
	BeforeLines  int        // Lines before match (-B)
	AfterLines   int        // Lines after match (-A)
	ContextLines int        // Lines before and after (-C)
	MaxResults   int        // Maximum results to return
	IgnoreCase   bool       // Case insensitive (-i)
	ShowLineNums bool       // Show line numbers (-n)
	IsRegex      bool       // Treat pattern as regex
	FilePattern  string     // Glob pattern for files (e.g., "*.md")
}

// GrepResult represents a search match
type GrepResult struct {
	File         string   `json:"file"`
	Project      string   `json:"project,omitempty"`
	Line         int      `json:"line,omitempty"`
	Column       int      `json:"column,omitempty"`
	Match        string   `json:"match,omitempty"`
	Context      []string `json:"context,omitempty"`
	Count        int      `json:"count,omitempty"`
}

// GrepResponse is the response format for grep operations
type GrepResponse struct {
	Matches      []GrepResult `json:"matches"`
	TotalMatches int          `json:"total_matches"`
	FilesMatched int          `json:"files_matched"`
	Mode         OutputMode   `json:"mode"`
}

// GlobOptions configures glob-like file finding
type GlobOptions struct {
	Pattern  string // Glob pattern (e.g., "**/*.md")
	Path     string // Base path to search from
	SortByTime bool // Sort by modification time
}

// GlobResult represents a found file
type GlobResult struct {
	Path    string `json:"path"`
	Project string `json:"project,omitempty"`
	Size    int64  `json:"size"`
	ModTime int64  `json:"mod_time"`
}

// SearchEngineV2 provides Claude Code-like search capabilities
type SearchEngineV2 struct {
	contentDir string
}

// NewSearchEngineV2 creates a new search engine
func NewSearchEngineV2(contentDir string) *SearchEngineV2 {
	return &SearchEngineV2{
		contentDir: contentDir,
	}
}

// Grep performs a grep-like search through documentation
func (e *SearchEngineV2) Grep(opts GrepOptions) (*GrepResponse, error) {
	response := &GrepResponse{
		Mode:    opts.OutputMode,
		Matches: []GrepResult{},
	}
	
	// Compile regex if needed
	var matcher func(string) []int
	if opts.IsRegex {
		re, err := regexp.Compile(opts.Pattern)
		if err != nil {
			return nil, fmt.Errorf("invalid regex: %w", err)
		}
		matcher = func(s string) []int {
			loc := re.FindStringIndex(s)
			if loc != nil {
				return loc
			}
			return nil
		}
	} else {
		pattern := opts.Pattern
		if opts.IgnoreCase {
			pattern = strings.ToLower(pattern)
		}
		matcher = func(s string) []int {
			searchIn := s
			if opts.IgnoreCase {
				searchIn = strings.ToLower(s)
			}
			idx := strings.Index(searchIn, pattern)
			if idx >= 0 {
				return []int{idx, idx + len(pattern)}
			}
			return nil
		}
	}
	
	// Determine search path
	searchPath := filepath.Join(e.contentDir, opts.Path)
	if opts.Path == "" {
		searchPath = e.contentDir
	}
	
	// Track files with matches
	filesWithMatches := make(map[string]bool)
	
	// Walk through files
	err := filepath.Walk(searchPath, func(path string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() {
			return nil
		}
		
		// Apply file pattern filter
		if opts.FilePattern != "" {
			matched, _ := filepath.Match(opts.FilePattern, filepath.Base(path))
			if !matched {
				return nil
			}
		}
		
		// Only search markdown files by default
		if !strings.HasSuffix(path, ".md") && opts.FilePattern == "" {
			return nil
		}
		
		// Search file
		matches, err := e.searchFile(path, matcher, opts)
		if err != nil {
			return nil
		}
		
		if len(matches) > 0 {
			filesWithMatches[path] = true
			
			// Get relative path and project
			relPath, _ := filepath.Rel(e.contentDir, path)
			parts := strings.Split(relPath, string(filepath.Separator))
			project := ""
			if len(parts) > 0 {
				project = parts[0]
			}
			
			// Add matches based on output mode
			switch opts.OutputMode {
			case OutputFiles:
				response.Matches = append(response.Matches, GrepResult{
					File:    relPath,
					Project: project,
				})
			case OutputContent:
				for _, match := range matches {
					match.File = relPath
					match.Project = project
					response.Matches = append(response.Matches, match)
				}
			case OutputCount:
				response.Matches = append(response.Matches, GrepResult{
					File:    relPath,
					Project: project,
					Count:   len(matches),
				})
			}
			
			response.TotalMatches += len(matches)
			
			// Limit results
			if opts.MaxResults > 0 && len(response.Matches) >= opts.MaxResults {
				return filepath.SkipDir
			}
		}
		
		return nil
	})
	
	if err != nil {
		return nil, err
	}
	
	response.FilesMatched = len(filesWithMatches)
	
	return response, nil
}

// searchFile searches within a single file
func (e *SearchEngineV2) searchFile(path string, matcher func(string) []int, opts GrepOptions) ([]GrepResult, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	
	var results []GrepResult
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
		
		loc := matcher(line)
		if loc != nil {
			result := GrepResult{
				Line:   lineNum,
				Column: loc[0] + 1,
				Match:  line,
			}
			
			// Get context if requested
			if opts.OutputMode == OutputContent {
				contextBefore := opts.BeforeLines
				contextAfter := opts.AfterLines
				if opts.ContextLines > 0 {
					contextBefore = opts.ContextLines
					contextAfter = opts.ContextLines
				}
				
				if contextBefore > 0 || contextAfter > 0 {
					var context []string
					
					// Before context
					start := maxInt(0, i-contextBefore)
					for j := start; j < i; j++ {
						context = append(context, lines[j])
					}
					
					// Current line
					context = append(context, line)
					
					// After context
					end := minInt(len(lines), i+contextAfter+1)
					for j := i + 1; j < end; j++ {
						context = append(context, lines[j])
					}
					
					result.Context = context
				}
			}
			
			results = append(results, result)
		}
	}
	
	return results, nil
}

// Glob finds files matching a pattern
func (e *SearchEngineV2) Glob(opts GlobOptions) ([]GlobResult, error) {
	var results []GlobResult
	
	searchPath := filepath.Join(e.contentDir, opts.Path)
	if opts.Path == "" {
		searchPath = e.contentDir
	}
	
	// Convert glob pattern to work with filepath.Walk
	err := filepath.Walk(searchPath, func(path string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() {
			return nil
		}
		
		relPath, _ := filepath.Rel(e.contentDir, path)
		
		// Match against pattern
		matched, err := filepath.Match(opts.Pattern, filepath.Base(path))
		if err != nil {
			return nil
		}
		
		// Also try matching the full relative path for patterns like "**/*.md"
		if !matched && strings.Contains(opts.Pattern, "/") {
			matched, _ = matchGlobPattern(opts.Pattern, relPath)
		}
		
		if matched {
			parts := strings.Split(relPath, string(filepath.Separator))
			project := ""
			if len(parts) > 0 {
				project = parts[0]
			}
			
			results = append(results, GlobResult{
				Path:    relPath,
				Project: project,
				Size:    info.Size(),
				ModTime: info.ModTime().Unix(),
			})
		}
		
		return nil
	})
	
	if err != nil {
		return nil, err
	}
	
	// Sort by modification time if requested
	if opts.SortByTime {
		sort.Slice(results, func(i, j int) bool {
			return results[i].ModTime > results[j].ModTime
		})
	}
	
	return results, nil
}

// matchGlobPattern matches a glob pattern with ** support
func matchGlobPattern(pattern, path string) (bool, error) {
	// Simple implementation - for full glob support, use a library
	if strings.Contains(pattern, "**") {
		// Convert ** to regex
		regexPattern := strings.ReplaceAll(pattern, "**", ".*")
		regexPattern = strings.ReplaceAll(regexPattern, "*", "[^/]*")
		regexPattern = "^" + regexPattern + "$"
		
		matched, err := regexp.MatchString(regexPattern, path)
		return matched, err
	}
	
	return filepath.Match(pattern, path)
}

// FormatResponse formats the response based on output format
func FormatResponse(resp *GrepResponse, format string) string {
	if format == "json" {
		data, _ := json.MarshalIndent(resp, "", "  ")
		return string(data)
	}
	
	// Human-readable format
	var output strings.Builder
	
	switch resp.Mode {
	case OutputFiles:
		output.WriteString(fmt.Sprintf("Found in %d files:\n\n", resp.FilesMatched))
		for _, match := range resp.Matches {
			output.WriteString(fmt.Sprintf("%s\n", match.File))
		}
		
	case OutputContent:
		output.WriteString(fmt.Sprintf("Found %d matches in %d files:\n\n", resp.TotalMatches, resp.FilesMatched))
		currentFile := ""
		for _, match := range resp.Matches {
			if match.File != currentFile {
				currentFile = match.File
				output.WriteString(fmt.Sprintf("\n%s:\n", match.File))
			}
			
			if match.Line > 0 {
				output.WriteString(fmt.Sprintf("  %d: %s\n", match.Line, match.Match))
			} else {
				output.WriteString(fmt.Sprintf("  %s\n", match.Match))
			}
			
			if len(match.Context) > 0 {
				for _, line := range match.Context {
					output.WriteString(fmt.Sprintf("     %s\n", line))
				}
			}
		}
		
	case OutputCount:
		output.WriteString(fmt.Sprintf("Match counts in %d files:\n\n", resp.FilesMatched))
		for _, match := range resp.Matches {
			output.WriteString(fmt.Sprintf("%s: %d matches\n", match.File, match.Count))
		}
	}
	
	return output.String()
}

// Use the helper functions from engine.go (maxInt and minInt)