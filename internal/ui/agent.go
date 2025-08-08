package ui

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/dorkitude/vd/internal/models"
	"github.com/dorkitude/vd/internal/scraper"
)

// BrowseResponse for agent mode
type BrowseResponse struct {
	Projects []ProjectInfo `json:"projects"`
	Total    int           `json:"total"`
}

type ProjectInfo struct {
	Name      string   `json:"name"`
	Path      string   `json:"path"`
	FileCount int      `json:"file_count"`
	Files     []string `json:"files,omitempty"`
	Status    string   `json:"status,omitempty"`
}

// RunBrowseAgent runs browse in agent mode
func RunBrowseAgent() {
	contentDir := "content"
	entries, err := ioutil.ReadDir(contentDir)
	if err != nil {
		fmt.Fprintf(os.Stderr, `{"error":"%s"}`, err.Error())
		os.Exit(1)
	}
	
	response := BrowseResponse{
		Projects: []ProjectInfo{},
	}
	
	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}
		
		projectPath := filepath.Join(contentDir, entry.Name())
		metadataPath := filepath.Join(projectPath, "vd.json")
		
		// Check if valid project
		if _, err := os.Stat(metadataPath); err != nil {
			continue
		}
		
		// Count markdown files
		files := []string{}
		filepath.Walk(projectPath, func(path string, info os.FileInfo, err error) error {
			if err == nil && strings.HasSuffix(path, ".md") {
				relPath, _ := filepath.Rel(projectPath, path)
				files = append(files, relPath)
			}
			return nil
		})
		
		// Get status from metadata
		status := "unknown"
		if data, err := ioutil.ReadFile(metadataPath); err == nil {
			var metadata models.Metadata
			if json.Unmarshal(data, &metadata) == nil {
				if s, ok := metadata.Metadata["status"].(string); ok {
					status = s
				}
			}
		}
		
		response.Projects = append(response.Projects, ProjectInfo{
			Name:      entry.Name(),
			Path:      projectPath,
			FileCount: len(files),
			Files:     files,
			Status:    status,
		})
	}
	
	response.Total = len(response.Projects)
	
	data, _ := json.MarshalIndent(response, "", "  ")
	fmt.Println(string(data))
}

// AddResponse for agent mode
type AddResponse struct {
	Success bool   `json:"success"`
	Project string `json:"project"`
	Message string `json:"message"`
	Path    string `json:"path,omitempty"`
}

// RunAddAgent runs add in agent mode
func RunAddAgent(args []string) {
	if len(args) < 2 {
		response := AddResponse{
			Success: false,
			Message: "Usage: vd add --agent <url> <name> [--type=mintlify|modal]",
		}
		data, _ := json.Marshal(response)
		fmt.Println(string(data))
		os.Exit(1)
	}
	
	url := args[0]
	name := args[1]
	docType := "mintlify"
	
	// Parse optional type flag
	for _, arg := range args[2:] {
		if strings.HasPrefix(arg, "--type=") {
			docType = strings.TrimPrefix(arg, "--type=")
		}
	}
	
	// Create project directory
	projectPath := filepath.Join("content", name)
	if err := os.MkdirAll(projectPath, 0755); err != nil {
		response := AddResponse{
			Success: false,
			Project: name,
			Message: fmt.Sprintf("Failed to create directory: %v", err),
		}
		data, _ := json.Marshal(response)
		fmt.Println(string(data))
		os.Exit(1)
	}
	
	// Create metadata
	metadata := models.Metadata{
		Title:       name,
		Description: fmt.Sprintf("Documentation from %s", url),
		RootURL:     url,
		Version:     "1.0.0",
		Files:       []models.FileMetadata{},
		Metadata: map[string]interface{}{
			"doc_type": docType,
			"status":   "pending_scrape",
		},
	}
	
	// Save metadata
	metadataPath := filepath.Join(projectPath, "vd.json")
	data, _ := json.MarshalIndent(metadata, "", "  ")
	if err := ioutil.WriteFile(metadataPath, data, 0644); err != nil {
		response := AddResponse{
			Success: false,
			Project: name,
			Message: fmt.Sprintf("Failed to save metadata: %v", err),
		}
		output, _ := json.Marshal(response)
		fmt.Println(string(output))
		os.Exit(1)
	}
	
	response := AddResponse{
		Success: true,
		Project: name,
		Message: fmt.Sprintf("Added %s documentation source", name),
		Path:    projectPath,
	}
	output, _ := json.Marshal(response)
	fmt.Println(string(output))
}

// ScrapeResponse for agent mode
type ScrapeResponse struct {
	Status          string            `json:"status"`
	ProjectsScraped []ScrapedProject  `json:"projects_scraped"`
	Errors          []string          `json:"errors,omitempty"`
}

type ScrapedProject struct {
	Name       string `json:"name"`
	FilesAdded int    `json:"files_added"`
	Status     string `json:"status"`
	Error      string `json:"error,omitempty"`
}

// RunScrapeAgent runs scrape in agent mode
func RunScrapeAgent(args []string) {
	response := ScrapeResponse{
		Status:          "completed",
		ProjectsScraped: []ScrapedProject{},
		Errors:          []string{},
	}
	
	// If specific project specified
	var projectFilter string
	if len(args) > 0 && !strings.HasPrefix(args[0], "--") {
		projectFilter = args[0]
	}
	
	contentDir := "content"
	entries, err := ioutil.ReadDir(contentDir)
	if err != nil {
		response.Status = "error"
		response.Errors = append(response.Errors, err.Error())
		data, _ := json.Marshal(response)
		fmt.Println(string(data))
		os.Exit(1)
	}
	
	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}
		
		// Apply filter if specified
		if projectFilter != "" && entry.Name() != projectFilter {
			continue
		}
		
		projectPath := filepath.Join(contentDir, entry.Name())
		metadataPath := filepath.Join(projectPath, "vd.json")
		
		// Read metadata
		data, err := ioutil.ReadFile(metadataPath)
		if err != nil {
			continue
		}
		
		var metadata models.Metadata
		if err := json.Unmarshal(data, &metadata); err != nil {
			continue
		}
		
		// Check if needs scraping
		status := "unknown"
		if s, ok := metadata.Metadata["status"].(string); ok {
			status = s
		}
		
		if status != "pending_scrape" && status != "unknown" {
			continue
		}
		
		// Determine scraper type
		docType := "mintlify"
		if dt, ok := metadata.Metadata["doc_type"].(string); ok {
			docType = dt
		}
		
		// Run appropriate scraper
		project := ScrapedProject{
			Name:   entry.Name(),
			Status: "success",
		}
		
		switch docType {
		case "modal":
			s := scraper.NewModalScraper(projectPath)
			if err := s.Scrape(); err != nil {
				project.Status = "error"
				project.Error = err.Error()
				response.Errors = append(response.Errors, fmt.Sprintf("%s: %v", entry.Name(), err))
			}
		default:
			// Use Python scraper for now
			project.Status = "skipped"
			project.Error = "Python scraper not available in agent mode"
		}
		
		// Count files added
		if project.Status == "success" {
			filepath.Walk(projectPath, func(path string, info os.FileInfo, err error) error {
				if err == nil && strings.HasSuffix(path, ".md") {
					project.FilesAdded++
				}
				return nil
			})
		}
		
		response.ProjectsScraped = append(response.ProjectsScraped, project)
	}
	
	if len(response.Errors) > 0 {
		response.Status = "completed_with_errors"
	}
	
	output, _ := json.MarshalIndent(response, "", "  ")
	fmt.Println(string(output))
}