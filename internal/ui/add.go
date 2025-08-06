package ui

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/dorkitude/vd/internal/models"
)

type AddModel struct {
	inputs     []textinput.Model
	focusIndex int
	err        error
	success    bool
	width      int
	height     int
}

const (
	folderNameIdx = iota
	titleIdx
	descriptionIdx
	urlIdx
)

func NewAddModel() *AddModel {
	m := &AddModel{
		inputs: make([]textinput.Model, 4),
	}

	var t textinput.Model
	for i := range m.inputs {
		t = textinput.New()
		t.Cursor.Blink = true
		t.CharLimit = 256

		switch i {
		case folderNameIdx:
			t.Placeholder = "folder-name (e.g., react-docs)"
			t.Focus()
			t.PromptStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("170"))
			t.TextStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("255"))
		case titleIdx:
			t.Placeholder = "Documentation Title"
		case descriptionIdx:
			t.Placeholder = "Brief description of the documentation"
		case urlIdx:
			t.Placeholder = "https://example.com/docs (Mintlify or GitHub URL)"
		}

		m.inputs[i] = t
	}

	return m
}

func (m *AddModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m *AddModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc":
			return m, tea.Quit
		case "tab", "shift+tab", "enter", "up", "down":
			s := msg.String()

			if s == "enter" && m.focusIndex == len(m.inputs)-1 {
				return m, m.createProject
			}

			if s == "up" || s == "shift+tab" {
				m.focusIndex--
			} else {
				m.focusIndex++
			}

			if m.focusIndex > len(m.inputs)-1 {
				m.focusIndex = 0
			} else if m.focusIndex < 0 {
				m.focusIndex = len(m.inputs) - 1
			}

			cmds := make([]tea.Cmd, len(m.inputs))
			for i := 0; i <= len(m.inputs)-1; i++ {
				if i == m.focusIndex {
					cmds[i] = m.inputs[i].Focus()
					m.inputs[i].PromptStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("170"))
					m.inputs[i].TextStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("255"))
				} else {
					m.inputs[i].Blur()
					m.inputs[i].PromptStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
					m.inputs[i].TextStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
				}
			}

			return m, tea.Batch(cmds...)
		}
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	}

	cmd := m.updateInputs(msg)

	return m, cmd
}

func (m *AddModel) updateInputs(msg tea.Msg) tea.Cmd {
	cmds := make([]tea.Cmd, len(m.inputs))

	for i := range m.inputs {
		m.inputs[i], cmds[i] = m.inputs[i].Update(msg)
	}

	return tea.Batch(cmds...)
}

func (m *AddModel) View() string {
	if m.success {
		return DocStyle.Render(
			SuccessStyle.Render("✅ Project created successfully!") + "\n\n" +
				"Folder: " + m.inputs[folderNameIdx].Value() + "\n" +
				"You can now run a scraper to populate the documentation.\n\n" +
				HelpStyle.Render("Press [esc] or [ctrl+c] to exit"),
		)
	}

	var s strings.Builder

	s.WriteString(TitleStyle.Render("➕ Add New Documentation Project"))
	s.WriteString("\n\n")

	labels := []string{
		"Folder Name:",
		"Title:",
		"Description:",
		"Source URL:",
	}

	for i := range m.inputs {
		s.WriteString(labels[i])
		s.WriteString("\n")
		s.WriteString(m.inputs[i].View())
		if i < len(m.inputs)-1 {
			s.WriteString("\n\n")
		}
	}

	s.WriteString("\n\n")
	
	if m.err != nil {
		s.WriteString(ErrorStyle.Render(fmt.Sprintf("Error: %v", m.err)))
		s.WriteString("\n\n")
	}

	button := "[ Submit ]"
	if m.focusIndex == len(m.inputs)-1 {
		button = SelectedItemStyle.Render("[ Submit ]")
	}
	s.WriteString(button)

	s.WriteString("\n\n")
	s.WriteString(HelpStyle.Render("tab/shift+tab to navigate • enter to submit • esc to cancel"))

	return DocStyle.Render(s.String())
}

func (m *AddModel) createProject() tea.Msg {
	folderName := strings.TrimSpace(m.inputs[folderNameIdx].Value())
	title := strings.TrimSpace(m.inputs[titleIdx].Value())
	description := strings.TrimSpace(m.inputs[descriptionIdx].Value())
	url := strings.TrimSpace(m.inputs[urlIdx].Value())

	if folderName == "" || title == "" || url == "" {
		m.err = fmt.Errorf("folder name, title, and URL are required")
		return nil
	}

	projectPath := filepath.Join("content", folderName)
	
	if _, err := os.Stat(projectPath); !os.IsNotExist(err) {
		m.err = fmt.Errorf("project folder already exists")
		return nil
	}

	if err := os.MkdirAll(projectPath, 0755); err != nil {
		m.err = err
		return nil
	}

	docType := "github"
	if strings.Contains(url, "mintlify") || strings.Contains(url, "docs") {
		docType = "mintlify"
	}

	metadata := models.Metadata{
		Title:       title,
		Description: description,
		RootURL:     url,
		ScrapeDate:  time.Now(),
		Version:     "1.0.0",
		Files:       []models.FileMetadata{},
		Metadata: map[string]interface{}{
			"doc_type": docType,
			"status":   "pending_scrape",
		},
	}

	data, err := json.MarshalIndent(metadata, "", "  ")
	if err != nil {
		m.err = err
		return nil
	}

	metadataPath := filepath.Join(projectPath, "vd.json")
	if err := ioutil.WriteFile(metadataPath, data, 0644); err != nil {
		m.err = err
		return nil
	}

	readmePath := filepath.Join(projectPath, "README.md")
	readmeContent := fmt.Sprintf("# %s\n\n%s\n\nSource: %s\n\n*Documentation will be populated after scraping.*",
		title, description, url)
	
	if err := ioutil.WriteFile(readmePath, []byte(readmeContent), 0644); err != nil {
		m.err = err
		return nil
	}

	m.success = true
	return nil
}