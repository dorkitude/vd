package ui

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/glamour"
	"github.com/dorkitude/vd/internal/search"
)

type SearchModel struct {
	project      string
	query        string
	textInput    textinput.Model
	results      []search.Result
	selectedIdx  int
	viewport     viewport.Model
	showingFile  bool
	fileContent  string
	searching    bool
	err          error
	width        int
	height       int
	engine       *search.Engine
}

func NewSearchModel(project, query string) *SearchModel {
	ti := textinput.New()
	ti.Placeholder = "Enter search query..."
	ti.Focus()
	ti.CharLimit = 256
	ti.Width = 50
	
	if query != "" {
		ti.SetValue(query)
	}
	
	vp := viewport.New(80, 20)
	
	return &SearchModel{
		project:   project,
		query:     query,
		textInput: ti,
		viewport:  vp,
		engine:    search.NewEngine("content"),
	}
}

func (m *SearchModel) Init() tea.Cmd {
	// If we have a query, search immediately
	if m.query != "" {
		return m.performSearch()
	}
	return textinput.Blink
}

func (m *SearchModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		m.viewport.Width = msg.Width
		m.viewport.Height = msg.Height - 6
		
	case tea.KeyMsg:
		if m.showingFile {
			switch msg.String() {
			case "q", "esc":
				m.showingFile = false
				return m, nil
			case "j", "down":
				m.viewport.LineDown(1)
			case "k", "up":
				m.viewport.LineUp(1)
			case "ctrl+d":
				m.viewport.HalfPageDown()
			case "ctrl+u":
				m.viewport.HalfPageUp()
			}
			
			var cmd tea.Cmd
			m.viewport, cmd = m.viewport.Update(msg)
			return m, cmd
		}
		
		switch msg.String() {
		case "ctrl+c", "esc":
			if m.query == "" {
				return m, tea.Quit
			}
			// Clear search and results
			m.query = ""
			m.textInput.SetValue("")
			m.results = nil
			return m, nil
			
		case "enter":
			if !m.searching {
				if m.results != nil && len(m.results) > 0 && m.selectedIdx >= 0 {
					// Open selected result
					return m, m.openResult()
				} else {
					// Perform search
					m.query = m.textInput.Value()
					return m, m.performSearch()
				}
			}
			
		case "tab", "shift+tab":
			// Switch between search box and results
			if m.results != nil && len(m.results) > 0 {
				m.textInput.Blur()
			} else {
				m.textInput.Focus()
			}
			
		case "up", "k":
			if m.results != nil && m.selectedIdx > 0 {
				m.selectedIdx--
			}
			
		case "down", "j":
			if m.results != nil && m.selectedIdx < len(m.results)-1 {
				m.selectedIdx++
			}
			
		case "ctrl+f":
			// Toggle fuzzy search
			// This would be implemented with a flag
			
		case "?":
			// Show help
			m.showingFile = true
			m.fileContent = searchHelpText
			m.viewport.SetContent(m.fileContent)
			return m, nil
		}
		
		if m.textInput.Focused() {
			var cmd tea.Cmd
			m.textInput, cmd = m.textInput.Update(msg)
			cmds = append(cmds, cmd)
		}
		
	case searchResultsMsg:
		m.searching = false
		m.results = msg.results
		m.err = msg.err
		if len(m.results) > 0 {
			m.selectedIdx = 0
		}
		return m, nil
	}
	
	return m, tea.Batch(cmds...)
}

func (m *SearchModel) View() string {
	if m.showingFile {
		return m.renderFileView()
	}
	
	var s strings.Builder
	
	// Title
	s.WriteString(TitleStyle.Render("🔍 Search Documentation"))
	s.WriteString("\n\n")
	
	// Search input
	s.WriteString("Query: ")
	s.WriteString(m.textInput.View())
	s.WriteString("\n\n")
	
	// Status
	if m.searching {
		s.WriteString("Searching...\n\n")
	} else if m.err != nil {
		s.WriteString(ErrorStyle.Render(fmt.Sprintf("Error: %v", m.err)))
		s.WriteString("\n\n")
	} else if m.results != nil {
		if len(m.results) == 0 {
			s.WriteString("No results found.\n\n")
		} else {
			s.WriteString(fmt.Sprintf("Found %d results:\n\n", len(m.results)))
			
			// Display results
			currentProject := ""
			for i, result := range m.results {
				// Group by project
				if result.Project != currentProject {
					currentProject = result.Project
					s.WriteString(fmt.Sprintf("\n📁 %s\n", HelpStyle.Render(result.Project)))
				}
				
				// Highlight selected
				resultStr := fmt.Sprintf("  %s:%d - %s",
					result.Path,
					result.Line,
					truncateString(result.Context, 60))
				
				if i == m.selectedIdx {
					s.WriteString("> " + SelectedItemStyle.Render(resultStr))
				} else {
					s.WriteString("  " + resultStr)
				}
				s.WriteString("\n")
				
				// Show max 20 results in UI
				if i >= 19 {
					remaining := len(m.results) - 20
					if remaining > 0 {
						s.WriteString(fmt.Sprintf("\n  ... and %d more results\n", remaining))
					}
					break
				}
			}
		}
	}
	
	// Help
	s.WriteString("\n")
	s.WriteString(HelpStyle.Render("[enter] search/open • [↑↓] navigate • [?] help • [esc] back"))
	
	return DocStyle.Render(s.String())
}

func (m *SearchModel) renderFileView() string {
	return fmt.Sprintf(
		"%s\n\n%s\n\n%s",
		TitleStyle.Render("📄 File Content"),
		m.viewport.View(),
		HelpStyle.Render("[q/esc] back • [j/k] scroll"),
	)
}

func (m *SearchModel) performSearch() tea.Cmd {
	return func() tea.Msg {
		m.searching = true
		
		opts := search.SearchOptions{
			Project:      m.project,
			Query:        m.query,
			ContextLines: 2,
			MaxResults:   100,
			FuzzyMatch:   false,
		}
		
		results, err := m.engine.Search(opts)
		
		return searchResultsMsg{
			results: results,
			err:     err,
		}
	}
}

func (m *SearchModel) openResult() tea.Cmd {
	if m.selectedIdx < 0 || m.selectedIdx >= len(m.results) {
		return nil
	}
	
	result := m.results[m.selectedIdx]
	filePath := fmt.Sprintf("content/%s/%s", result.Project, result.Path)
	
	return func() tea.Msg {
		content, err := os.ReadFile(filePath)
		if err != nil {
			m.err = err
			return nil
		}
		
		// Render with glamour
		renderer, _ := glamour.NewTermRenderer(
			glamour.WithAutoStyle(),
			glamour.WithWordWrap(m.width-4),
		)
		
		rendered, err := renderer.Render(string(content))
		if err != nil {
			rendered = string(content)
		}
		
		m.fileContent = rendered
		m.viewport.SetContent(rendered)
		m.showingFile = true
		
		// Try to scroll to the line
		if result.Line > 0 {
			// Approximate line position
			m.viewport.SetYOffset(result.Line - 5)
		}
		
		return nil
	}
}

type searchResultsMsg struct {
	results []search.Result
	err     error
}

// RunSearchCLI runs search from command line
func RunSearchCLI(project, query string) {
	engine := search.NewEngine("content")
	
	// Check for environment variables
	format := os.Getenv("VD_FORMAT")
	contextLines := 2
	if ctx := os.Getenv("VD_CONTEXT"); ctx != "" {
		if n, err := strconv.Atoi(ctx); err == nil {
			contextLines = n / 50 // Convert characters to approximate lines
		}
	}
	
	opts := search.SearchOptions{
		Project:      project,
		Query:        query,
		ContextLines: contextLines,
		MaxResults:   50,
		FuzzyMatch:   false,
		OutputFormat: format,
	}
	
	results, err := engine.Search(opts)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	
	// Format and output results
	output := search.FormatResults(results, format)
	fmt.Print(output)
}

func truncateString(s string, maxLen int) string {
	s = strings.TrimSpace(s)
	if len(s) <= maxLen {
		return s
	}
	return s[:maxLen-3] + "..."
}

const searchHelpText = `# Search Help

## Commands
- **Enter**: Perform search or open selected result
- **↑/↓ or j/k**: Navigate through results
- **Tab**: Switch focus between search box and results
- **Esc**: Clear search or go back
- **?**: Show this help
- **q**: Quit

## Search Tips
- Search is case-insensitive by default
- Searches through all markdown files in documentation
- Results are sorted by relevance score

## Command Line Usage

### Search all projects:
` + "```" + `bash
vd search "your query"
` + "```" + `

### Search specific project:
` + "```" + `bash
vd search project_name "your query"
` + "```" + `

### JSON output for AI agents:
` + "```" + `bash
VD_FORMAT=json vd search "query"
` + "```" + `

### Adjust context size:
` + "```" + `bash
VD_CONTEXT=200 vd search "query"
` + "```" + ``