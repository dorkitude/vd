package scraper

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/dorkitude/vd/internal/models"
)

type ModalScraper struct {
	baseURL     string
	projectPath string
	client      *http.Client
	scraped     map[string]bool
}

func NewModalScraper(projectPath string) *ModalScraper {
	return &ModalScraper{
		baseURL:     "https://modal.com/docs",
		projectPath: projectPath,
		client: &http.Client{
			Timeout: 30 * time.Second,
		},
		scraped: make(map[string]bool),
	}
}

func (s *ModalScraper) Scrape() error {
	fmt.Println("🕷️  Starting Modal docs scrape...")
	
	// Create project directory
	if err := os.MkdirAll(s.projectPath, 0755); err != nil {
		return fmt.Errorf("failed to create project directory: %w", err)
	}

	// Essential Modal docs sections
	sections := []string{
		"",  // Main docs page
		"/guide",
		"/guide/environments", 
		"/guide/cloud-storage",
		"/guide/functions",
		"/guide/images",
		"/guide/webhooks",
		"/guide/secrets",
		"/guide/volumes",
		"/guide/dictionaries",
		"/examples",
		"/reference",
	}

	// Create metadata
	metadata := &models.Metadata{
		Title:       "Modal",
		Description: "Documentation scraped from https://modal.com/docs",
		RootURL:     s.baseURL,
		ScrapeDate:  time.Now(),
		Version:     "1.0.0",
		Files:       []models.FileMetadata{},
		Metadata: map[string]interface{}{
			"doc_type": "modal",
			"status":   "scraped",
		},
	}

	// Discover more pages from the main sections
	allPages := s.discoverPages(sections)
	fmt.Printf("📄 Found %d pages to scrape\n", len(allPages))

	// Scrape each page
	for i, pagePath := range allPages {
		fmt.Printf("[%d/%d] Scraping %s...\n", i+1, len(allPages), pagePath)
		
		if err := s.scrapePage(pagePath, metadata); err != nil {
			fmt.Printf("  ⚠️  Warning: Failed to scrape %s: %v\n", pagePath, err)
			continue
		}
		
		// Rate limiting
		time.Sleep(300 * time.Millisecond)
	}

	// Save metadata
	if err := s.saveMetadata(metadata); err != nil {
		return fmt.Errorf("failed to save metadata: %w", err)
	}

	fmt.Printf("✅ Successfully scraped %d files!\n", len(metadata.Files))
	return nil
}

func (s *ModalScraper) discoverPages(sections []string) []string {
	pages := make(map[string]bool)
	
	for _, section := range sections {
		url := s.baseURL + section
		pages[section] = true
		
		// Fetch the page and look for more doc links
		resp, err := s.client.Get(url)
		if err != nil {
			continue
		}
		defer resp.Body.Close()
		
		if resp.StatusCode != 200 {
			continue
		}
		
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			continue
		}
		
		// Find all /docs/ links
		re := regexp.MustCompile(`href="/docs/([^"#]+)"`)
		matches := re.FindAllStringSubmatch(string(body), -1)
		
		for _, match := range matches {
			if len(match) > 1 {
				path := "/" + match[1]
				// Skip non-documentation paths
				if !strings.Contains(path, "api/") && 
				   !strings.Contains(path, "auth/") &&
				   !strings.Contains(path, ".json") &&
				   !strings.Contains(path, ".xml") {
					pages[path] = true
				}
			}
		}
	}
	
	// Convert map to slice
	result := make([]string, 0, len(pages))
	for page := range pages {
		result = append(result, page)
	}
	
	// Limit to prevent excessive scraping
	if len(result) > 100 {
		result = result[:100]
	}
	
	return result
}

func (s *ModalScraper) scrapePage(pagePath string, metadata *models.Metadata) error {
	// Skip if already scraped
	if s.scraped[pagePath] {
		return nil
	}
	s.scraped[pagePath] = true
	
	// Build the URL
	pageURL := s.baseURL + pagePath
	
	resp, err := s.client.Get(pageURL)
	if err != nil {
		return fmt.Errorf("failed to fetch page: %w", err)
	}
	defer resp.Body.Close()
	
	if resp.StatusCode != 200 {
		return fmt.Errorf("got status %d", resp.StatusCode)
	}
	
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read body: %w", err)
	}
	
	// Extract content (simplified - ideally we'd use a proper HTML parser)
	content := s.extractContent(string(body), pagePath)
	
	// Create file path
	localPath := strings.TrimPrefix(pagePath, "/")
	if localPath == "" {
		localPath = "index"
	}
	if !strings.HasSuffix(localPath, ".md") {
		localPath += ".md"
	}
	
	// Ensure directory exists
	fullPath := filepath.Join(s.projectPath, localPath)
	dir := filepath.Dir(fullPath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}
	
	// Save the file
	if err := os.WriteFile(fullPath, []byte(content), 0644); err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}
	
	// Add to metadata
	metadata.Files = append(metadata.Files, models.FileMetadata{
		Path:      localPath,
		URL:       pageURL,
		Title:     s.extractTitle(string(body), pagePath),
		ScrapedAt: time.Now(),
	})
	
	return nil
}

func (s *ModalScraper) extractContent(html string, pagePath string) string {
	// This is a simplified extraction
	// In production, we'd use a proper HTML to Markdown converter
	
	title := s.extractTitle(html, pagePath)
	
	// For now, create a basic markdown file with a note
	content := fmt.Sprintf(`# %s

> Note: This is a placeholder for Modal documentation from %s%s
> The full HTML-to-Markdown conversion will be implemented with a proper HTML parser.

## About This Page

This documentation was scraped from the Modal docs website. Modal is a cloud platform for running Python code in the cloud with minimal configuration.

Visit the original documentation at: %s%s
`, title, s.baseURL, pagePath, s.baseURL, pagePath)
	
	// Try to extract some basic content
	// Look for main content area
	mainStart := strings.Index(html, "<main")
	if mainStart > 0 {
		mainEnd := strings.Index(html[mainStart:], "</main>")
		if mainEnd > 0 {
			mainContent := html[mainStart : mainStart+mainEnd]
			// Extract text from paragraphs
			re := regexp.MustCompile(`<p[^>]*>([^<]+)</p>`)
			matches := re.FindAllStringSubmatch(mainContent, 5)
			if len(matches) > 0 {
				content += "\n## Content Preview\n\n"
				for _, match := range matches {
					if len(match) > 1 {
						content += match[1] + "\n\n"
					}
				}
			}
		}
	}
	
	return content
}

func (s *ModalScraper) extractTitle(html string, pagePath string) string {
	// Try to find h1
	re := regexp.MustCompile(`<h1[^>]*>([^<]+)</h1>`)
	matches := re.FindStringSubmatch(html)
	if len(matches) > 1 {
		return strings.TrimSpace(matches[1])
	}
	
	// Try title tag
	re = regexp.MustCompile(`<title>([^<]+)</title>`)
	matches = re.FindStringSubmatch(html)
	if len(matches) > 1 {
		title := strings.TrimSpace(matches[1])
		// Remove common suffixes
		title = strings.ReplaceAll(title, " | Modal", "")
		title = strings.ReplaceAll(title, " - Modal", "")
		return title
	}
	
	// Fallback to path
	parts := strings.Split(strings.Trim(pagePath, "/"), "/")
	if len(parts) > 0 {
		lastPart := parts[len(parts)-1]
		lastPart = strings.ReplaceAll(lastPart, "-", " ")
		lastPart = strings.ReplaceAll(lastPart, "_", " ")
		return strings.Title(lastPart)
	}
	
	return "Modal Documentation"
}

func (s *ModalScraper) saveMetadata(metadata *models.Metadata) error {
	data, err := json.MarshalIndent(metadata, "", "  ")
	if err != nil {
		return err
	}
	
	metadataPath := filepath.Join(s.projectPath, "vd.json")
	return os.WriteFile(metadataPath, data, 0644)
}