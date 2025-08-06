package ui

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/dorkitude/vd/internal/models"
)

type ScrapeModel struct {
	projects     []ScrapeProjectItem
	list         list.Model
	selected     *ScrapeProjectItem
	scraping     bool
	spinner      spinner.Model
	status       string
	err          error
	width        int
	height       int
}

type ScrapeProjectItem struct {
	name     string
	path     string
	metadata *models.Metadata
}

func (i ScrapeProjectItem) Title() string       { return i.name }
func (i ScrapeProjectItem) Description() string {
	if i.metadata != nil {
		status := "unknown"
		if s, ok := i.metadata.Metadata["status"].(string); ok {
			status = s
		}
		return fmt.Sprintf("%s [%s]", i.metadata.Description, status)
	}
	return "No metadata"
}
func (i ScrapeProjectItem) FilterValue() string { return i.name }

func NewScrapeModel() *ScrapeModel {
	projects := loadPendingProjects()
	
	items := make([]list.Item, len(projects))
	for i, p := range projects {
		items[i] = p
	}

	l := list.New(items, list.NewDefaultDelegate(), 80, 14)
	l.Title = "🕷️  Scrape Documentation"
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(true)
	l.Styles.Title = TitleStyle
	
	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("170"))

	return &ScrapeModel{
		projects: projects,
		list:     l,
		spinner:  s,
	}
}

func loadPendingProjects() []ScrapeProjectItem {
	var projects []ScrapeProjectItem
	
	contentDir := "content"
	entries, err := ioutil.ReadDir(contentDir)
	if err != nil {
		return projects
	}

	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}

		projectPath := filepath.Join(contentDir, entry.Name())
		metadataPath := filepath.Join(projectPath, "vd.json")
		
		data, err := ioutil.ReadFile(metadataPath)
		if err != nil {
			continue
		}

		var metadata models.Metadata
		if err := json.Unmarshal(data, &metadata); err != nil {
			continue
		}

		// Only show projects that need scraping
		status := "unknown"
		if s, ok := metadata.Metadata["status"].(string); ok {
			status = s
		}
		
		if status == "pending_scrape" || status == "unknown" {
			projects = append(projects, ScrapeProjectItem{
				name:     entry.Name(),
				path:     projectPath,
				metadata: &metadata,
			})
		}
	}

	return projects
}

func (m *ScrapeModel) Init() tea.Cmd {
	return nil
}

func (m *ScrapeModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		m.list.SetWidth(msg.Width)
		m.list.SetHeight(msg.Height - 4)

	case tea.KeyMsg:
		if m.scraping {
			// Can't do anything while scraping except quit
			if msg.String() == "ctrl+c" {
				return m, tea.Quit
			}
			return m, nil
		}

		switch msg.String() {
		case "q", "ctrl+c", "esc":
			return m, tea.Quit
		case "enter":
			if i, ok := m.list.SelectedItem().(ScrapeProjectItem); ok {
				m.selected = &i
				m.scraping = true
				m.status = "Starting scrape..."
				return m, tea.Batch(
					m.spinner.Tick,
					m.startScrape(),
				)
			}
		}

	case spinner.TickMsg:
		if m.scraping {
			var cmd tea.Cmd
			m.spinner, cmd = m.spinner.Update(msg)
			return m, cmd
		}

	case scrapeCompleteMsg:
		m.scraping = false
		if msg.err != nil {
			m.err = msg.err
			m.status = fmt.Sprintf("❌ Error: %v", msg.err)
		} else {
			m.status = fmt.Sprintf("✅ Successfully scraped %d files!", msg.fileCount)
			// Reload projects list
			m.projects = loadPendingProjects()
			items := make([]list.Item, len(m.projects))
			for i, p := range m.projects {
				items[i] = p
			}
			m.list.SetItems(items)
		}
		return m, nil

	case scrapeStatusMsg:
		m.status = string(msg)
		return m, nil
	}

	if !m.scraping {
		var cmd tea.Cmd
		m.list, cmd = m.list.Update(msg)
		return m, cmd
	}

	return m, nil
}

func (m *ScrapeModel) View() string {
	if m.scraping {
		return m.renderScraping()
	}

	s := m.list.View()
	s += "\n\n"
	
	if m.status != "" {
		s += m.status + "\n\n"
	}
	
	s += HelpStyle.Render("Press [enter] to scrape • [q] to quit")
	
	return DocStyle.Render(s)
}

func (m *ScrapeModel) renderScraping() string {
	var s strings.Builder
	
	s.WriteString(TitleStyle.Render("🕷️  Scraping Documentation"))
	s.WriteString("\n\n")
	
	if m.selected != nil && m.selected.metadata != nil {
		s.WriteString(fmt.Sprintf("Project: %s\n", m.selected.metadata.Title))
		s.WriteString(fmt.Sprintf("URL: %s\n", m.selected.metadata.RootURL))
		s.WriteString("\n")
	}
	
	s.WriteString(fmt.Sprintf("%s %s\n", m.spinner.View(), m.status))
	s.WriteString("\n")
	s.WriteString(HelpStyle.Render("Press [ctrl+c] to cancel"))
	
	return DocStyle.Render(s.String())
}

type scrapeCompleteMsg struct {
	err       error
	fileCount int
}

type scrapeStatusMsg string

func (m *ScrapeModel) startScrape() tea.Cmd {
	return func() tea.Msg {
		if m.selected == nil || m.selected.metadata == nil {
			return scrapeCompleteMsg{err: fmt.Errorf("no project selected")}
		}

		// Check if Python script exists
		scriptPath := "scripts/scrape_mintlify.py"
		if _, err := os.Stat(scriptPath); os.IsNotExist(err) {
			return scrapeCompleteMsg{err: fmt.Errorf("scraper script not found at %s", scriptPath)}
		}

		// Run the Python scraper
		cmd := exec.Command("python3", scriptPath, m.selected.metadata.RootURL, m.selected.path)
		
		output, err := cmd.CombinedOutput()
		if err != nil {
			return scrapeCompleteMsg{err: fmt.Errorf("scraper failed: %v\nOutput: %s", err, string(output))}
		}

		// Read the updated metadata to get file count
		metadataPath := filepath.Join(m.selected.path, "vd.json")
		data, err := ioutil.ReadFile(metadataPath)
		if err != nil {
			return scrapeCompleteMsg{err: fmt.Errorf("failed to read metadata: %v", err)}
		}

		var metadata models.Metadata
		if err := json.Unmarshal(data, &metadata); err != nil {
			return scrapeCompleteMsg{err: fmt.Errorf("failed to parse metadata: %v", err)}
		}

		return scrapeCompleteMsg{fileCount: len(metadata.Files)}
	}
}