package models

import "time"

type Metadata struct {
	Title       string            `json:"title"`
	Description string            `json:"description"`
	RootURL     string            `json:"root_url"`
	ScrapeDate  time.Time         `json:"scrape_date"`
	Version     string            `json:"version"`
	Files       []FileMetadata    `json:"files"`
	Metadata    map[string]interface{} `json:"metadata"`
}

type FileMetadata struct {
	Path      string    `json:"path"`
	URL       string    `json:"url"`
	Title     string    `json:"title"`
	ScrapedAt time.Time `json:"scraped_at"`
}