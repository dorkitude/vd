package ui

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/lipgloss"
	"github.com/dorkitude/vd/internal/models"
)

type BrowseModel struct {
	list         list.Model
	projects     []ProjectItem
	width        int
	height       int
	showDetails  bool
	showFile     bool
	selected     *ProjectItem
	fileList     list.Model
	fileViewer   viewport.Model
	currentFile  string
	viewMode     string // "list", "details", "files", "viewer"
}

type ProjectItem struct {
	name        string
	metadata    *models.Metadata
	path        string
}

type FileItem struct {
	path  string
	title string
	url   string
}

func (i ProjectItem) Title() string       { return i.name }
func (i ProjectItem) Description() string { 
	if i.metadata != nil {
		return i.metadata.Description
	}
	return "No description available"
}
func (i ProjectItem) FilterValue() string { return i.name }

func (i FileItem) Title() string       { return i.title }
func (i FileItem) Description() string { return i.path }
func (i FileItem) FilterValue() string { return i.title + " " + i.path }

func NewBrowseModel() *BrowseModel {
	projects := loadProjects()
	
	items := make([]list.Item, len(projects))
	for i, p := range projects {
		items[i] = p
	}

	const defaultWidth = 80
	const listHeight = 14

	l := list.New(items, list.NewDefaultDelegate(), defaultWidth, listHeight)
	l.Title = "📚 VD - Browse Documentation"
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(true)
	l.Styles.Title = TitleStyle
	l.Styles.PaginationStyle = HelpStyle
	l.Styles.HelpStyle = HelpStyle

	vp := viewport.New(defaultWidth, 20)
	vp.Style = lipgloss.NewStyle().
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("62")).
		PaddingRight(1)

	return &BrowseModel{
		list:       l,
		projects:   projects,
		fileViewer: vp,
		viewMode:   "list",
	}
}

func loadProjects() []ProjectItem {
	var projects []ProjectItem
	
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
		
		item := ProjectItem{
			name: entry.Name(),
			path: projectPath,
		}

		if data, err := ioutil.ReadFile(metadataPath); err == nil {
			var metadata models.Metadata
			if err := json.Unmarshal(data, &metadata); err == nil {
				item.metadata = &metadata
			}
		}

		projects = append(projects, item)
	}

	return projects
}

func (m *BrowseModel) Init() tea.Cmd {
	return nil
}

func (m *BrowseModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		m.list.SetWidth(msg.Width)
		m.list.SetHeight(msg.Height - 4)
		m.fileViewer.Width = msg.Width - 4
		m.fileViewer.Height = msg.Height - 6
		if m.fileList.Width() != 0 {
			m.fileList.SetWidth(msg.Width)
			m.fileList.SetHeight(msg.Height - 4)
		}
		return m, nil

	case tea.KeyMsg:
		switch m.viewMode {
		case "list":
			return m.updateProjectList(msg)
		case "details":
			return m.updateDetails(msg)
		case "files":
			return m.updateFileList(msg)
		case "viewer":
			return m.updateFileViewer(msg)
		}
	}

	return m, tea.Batch(cmds...)
}

func (m *BrowseModel) updateProjectList(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "q", "ctrl+c":
		return m, tea.Quit
	case "enter":
		if i, ok := m.list.SelectedItem().(ProjectItem); ok {
			m.selected = &i
			m.viewMode = "details"
		}
		return m, nil
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m *BrowseModel) updateDetails(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "q", "ctrl+c":
		return m, tea.Quit
	case "esc":
		m.viewMode = "list"
		m.selected = nil
		return m, nil
	case "f", "enter":
		// Show files list
		if m.selected != nil && m.selected.metadata != nil && len(m.selected.metadata.Files) > 0 {
			m.viewMode = "files"
			m.setupFileList()
		}
		return m, nil
	}
	return m, nil
}

func (m *BrowseModel) setupFileList() {
	if m.selected == nil || m.selected.metadata == nil {
		return
	}

	items := make([]list.Item, len(m.selected.metadata.Files))
	for i, f := range m.selected.metadata.Files {
		items[i] = FileItem{
			path:  f.Path,
			title: f.Title,
			url:   f.URL,
		}
	}

	const defaultWidth = 80
	const listHeight = 14

	m.fileList = list.New(items, list.NewDefaultDelegate(), defaultWidth, listHeight)
	m.fileList.Title = fmt.Sprintf("📄 Files in %s", m.selected.metadata.Title)
	m.fileList.SetShowStatusBar(false)
	m.fileList.SetFilteringEnabled(true)
	m.fileList.Styles.Title = TitleStyle
	m.fileList.Styles.PaginationStyle = HelpStyle
	m.fileList.Styles.HelpStyle = HelpStyle
	
	if m.width > 0 {
		m.fileList.SetWidth(m.width)
		m.fileList.SetHeight(m.height - 4)
	}
}

func (m *BrowseModel) updateFileList(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "q", "ctrl+c":
		return m, tea.Quit
	case "esc":
		m.viewMode = "details"
		return m, nil
	case "enter":
		if i, ok := m.fileList.SelectedItem().(FileItem); ok {
			m.currentFile = i.path
			m.viewMode = "viewer"
			return m, m.loadAndRenderFile(i.path)
		}
	}

	var cmd tea.Cmd
	m.fileList, cmd = m.fileList.Update(msg)
	return m, cmd
}

func (m *BrowseModel) updateFileViewer(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "q", "ctrl+c":
		return m, tea.Quit
	case "esc":
		m.viewMode = "files"
		return m, nil
	}

	var cmd tea.Cmd
	m.fileViewer, cmd = m.fileViewer.Update(msg)
	return m, cmd
}

func (m *BrowseModel) loadAndRenderFile(relativePath string) tea.Cmd {
	return func() tea.Msg {
		if m.selected == nil {
			return nil
		}

		fullPath := filepath.Join(m.selected.path, relativePath)
		content, err := ioutil.ReadFile(fullPath)
		if err != nil {
			m.fileViewer.SetContent(fmt.Sprintf("Error loading file: %v", err))
			return nil
		}

		// Use glamour to render markdown
		renderer, err := glamour.NewTermRenderer(
			glamour.WithAutoStyle(),
			glamour.WithWordWrap(m.width-6),
		)
		if err != nil {
			m.fileViewer.SetContent(string(content))
			return nil
		}

		rendered, err := renderer.Render(string(content))
		if err != nil {
			m.fileViewer.SetContent(string(content))
			return nil
		}

		m.fileViewer.SetContent(rendered)
		m.fileViewer.GotoTop()
		return nil
	}
}

func (m *BrowseModel) View() string {
	switch m.viewMode {
	case "details":
		return m.renderDetails()
	case "files":
		return DocStyle.Render(m.fileList.View() + "\n" + m.helpViewFiles())
	case "viewer":
		return m.renderFileViewer()
	default:
		return DocStyle.Render(m.list.View() + "\n" + m.helpView())
	}
}

func (m *BrowseModel) renderDetails() string {
	if m.selected == nil || m.selected.metadata == nil {
		return "No details available"
	}

	meta := m.selected.metadata
	
	var s strings.Builder
	
	s.WriteString(TitleStyle.Render(fmt.Sprintf("📁 %s", meta.Title)))
	s.WriteString("\n\n")
	
	info := fmt.Sprintf("📝 Description: %s\n🌐 Source: %s\n📅 Scraped: %s\n📚 Files: %d",
		meta.Description,
		meta.RootURL,
		meta.ScrapeDate.Format("2006-01-02"),
		len(meta.Files),
	)
	
	s.WriteString(InfoStyle.Render(info))
	s.WriteString("\n\n")
	
	if len(meta.Files) > 0 {
		s.WriteString(TitleStyle.Copy().UnsetMargins().Render("Files Preview:"))
		s.WriteString("\n")
		
		maxShow := 10
		for i, file := range meta.Files {
			if i >= maxShow {
				s.WriteString(fmt.Sprintf("\n  ... and %d more files", len(meta.Files)-maxShow))
				break
			}
			s.WriteString(fmt.Sprintf("\n  • %s", file.Path))
		}
		s.WriteString("\n\n")
		s.WriteString(SuccessStyle.Render("Press [f] or [enter] to browse files"))
	}
	
	s.WriteString("\n\n")
	s.WriteString(HelpStyle.Render("Press [esc] to go back • [f] to view files • [q] to quit"))
	
	return DocStyle.Render(s.String())
}

func (m *BrowseModel) renderFileViewer() string {
	header := TitleStyle.Render(fmt.Sprintf("📄 %s", m.currentFile))
	help := HelpStyle.Render("↑/↓ to scroll • [esc] back to files • [q] to quit")
	
	return DocStyle.Render(
		lipgloss.JoinVertical(
			lipgloss.Left,
			header,
			m.fileViewer.View(),
			help,
		),
	)
}

func (m *BrowseModel) helpView() string {
	return HelpStyle.Render("Press [enter] to view details • [/] to search • [q] to quit")
}

func (m *BrowseModel) helpViewFiles() string {
	return HelpStyle.Render("Press [enter] to view file • [/] to search • [esc] to go back • [q] to quit")
}