package scraper

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"

	"github.com/dorkitude/vd/internal/models"
)

type MintlifyScraper struct {
	baseURL     string
	projectPath string
	client      *http.Client
}

type MintConfig struct {
	Name        string         `json:"name"`
	Navigation  []NavItem      `json:"navigation"`
	Anchors     []Anchor       `json:"anchors"`
}

type NavItem struct {
	Group string   `json:"group"`
	Pages []string `json:"pages"`
}

type Anchor struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

func NewMintlifyScraper(baseURL, projectPath string) *MintlifyScraper {
	return &MintlifyScraper{
		baseURL:     strings.TrimSuffix(baseURL, "/"),
		projectPath: projectPath,
		client: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

func (s *MintlifyScraper) Scrape() error {
	fmt.Println("🔍 Fetching mint.json configuration...")
	
	// Try to fetch mint.json
	mintConfig, err := s.fetchMintConfig()
	if err != nil {
		return fmt.Errorf("failed to fetch mint.json: %w", err)
	}

	fmt.Printf("📚 Found documentation: %s\n", mintConfig.Name)
	
	// Create project directory
	if err := os.MkdirAll(s.projectPath, 0755); err != nil {
		return fmt.Errorf("failed to create project directory: %w", err)
	}

	// Collect all pages to scrape
	var pages []string
	for _, nav := range mintConfig.Navigation {
		pages = append(pages, nav.Pages...)
	}

	fmt.Printf("📄 Found %d pages to scrape\n", len(pages))

	// Create metadata
	metadata := &models.Metadata{
		Title:       mintConfig.Name,
		Description: fmt.Sprintf("Documentation scraped from %s", s.baseURL),
		RootURL:     s.baseURL,
		ScrapeDate:  time.Now(),
		Version:     "1.0.0",
		Files:       []models.FileMetadata{},
		Metadata: map[string]interface{}{
			"doc_type": "mintlify",
			"status":   "scraped",
		},
	}

	// Scrape each page
	for i, pagePath := range pages {
		fmt.Printf("[%d/%d] Scraping %s...\n", i+1, len(pages), pagePath)
		
		if err := s.scrapePage(pagePath, metadata); err != nil {
			fmt.Printf("  ⚠️  Warning: Failed to scrape %s: %v\n", pagePath, err)
			continue
		}
		
		// Be nice to the server
		time.Sleep(500 * time.Millisecond)
	}

	// Save metadata
	if err := s.saveMetadata(metadata); err != nil {
		return fmt.Errorf("failed to save metadata: %w", err)
	}

	fmt.Printf("✅ Successfully scraped %d files!\n", len(metadata.Files))
	return nil
}

func (s *MintlifyScraper) fetchMintConfig() (*MintConfig, error) {
	// Try common locations for mint.json
	urls := []string{
		s.baseURL + "/mint.json",
		s.baseURL + "/_next/data/mint.json",
		s.baseURL + "/api/mint.json",
	}

	for _, url := range urls {
		resp, err := s.client.Get(url)
		if err != nil {
			continue
		}
		defer resp.Body.Close()

		if resp.StatusCode == 200 {
			var config MintConfig
			if err := json.NewDecoder(resp.Body).Decode(&config); err != nil {
				continue
			}
			return &config, nil
		}
	}

	// If we can't find mint.json, create a basic config
	return &MintConfig{
		Name: "Documentation",
		Navigation: []NavItem{
			{
				Group: "Getting Started",
				Pages: []string{"introduction", "quickstart", "installation"},
			},
		},
	}, nil
}

func (s *MintlifyScraper) scrapePage(pagePath string, metadata *models.Metadata) error {
	// Build the URL
	pageURL := s.baseURL + "/" + strings.TrimPrefix(pagePath, "/")
	
	// For Mintlify, we need to fetch the raw markdown
	// Try different patterns
	rawURLs := []string{
		pageURL + ".mdx",
		pageURL + ".md",
		pageURL,
	}

	var content string
	var foundURL string
	
	for _, url := range rawURLs {
		resp, err := s.client.Get(url)
		if err != nil {
			continue
		}
		defer resp.Body.Close()

		if resp.StatusCode == 200 {
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				continue
			}
			
			// Check if it's HTML or Markdown
			bodyStr := string(body)
			if strings.Contains(bodyStr, "<!DOCTYPE") || strings.Contains(bodyStr, "<html") {
				// It's HTML, try to extract content
				content = s.extractMarkdownFromHTML(bodyStr)
			} else {
				// Assume it's markdown/mdx
				content = bodyStr
			}
			
			if content != "" {
				foundURL = url
				break
			}
		}
	}

	if content == "" {
		return fmt.Errorf("could not fetch content for %s", pagePath)
	}

	// Create file path
	localPath := pagePath + ".md"
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
		URL:       foundURL,
		Title:     s.extractTitle(content, pagePath),
		ScrapedAt: time.Now(),
	})

	return nil
}

func (s *MintlifyScraper) extractMarkdownFromHTML(html string) string {
	// This is a simple extraction - in production you'd use a proper HTML parser
	// For now, let's just return a placeholder
	return "# Page Content\n\nThis page was scraped from HTML. The raw content extraction is not yet implemented.\n"
}

func (s *MintlifyScraper) extractTitle(content, pagePath string) string {
	// Try to extract title from markdown
	lines := strings.Split(content, "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "# ") {
			return strings.TrimPrefix(line, "# ")
		}
	}
	
	// Fallback to path
	base := path.Base(pagePath)
	base = strings.ReplaceAll(base, "-", " ")
	base = strings.ReplaceAll(base, "_", " ")
	return strings.Title(base)
}

func (s *MintlifyScraper) saveMetadata(metadata *models.Metadata) error {
	data, err := json.MarshalIndent(metadata, "", "  ")
	if err != nil {
		return err
	}

	metadataPath := filepath.Join(s.projectPath, "vd.json")
	return os.WriteFile(metadataPath, data, 0644)
}