package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/dorkitude/vd/internal/docs"
	"github.com/dorkitude/vd/internal/ui"
)

func main() {
	// Check for global --agent flag or environment variable
	agentMode := os.Getenv("VD_AGENT") == "true" || hasFlag("--agent")
	
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "browse":
			runBrowse(agentMode)
		case "add":
			runAdd(agentMode)
		case "scrape":
			runScrape(agentMode)
		case "search", "grep":
			runSearch()
		case "glob", "find":
			runGlob()
		case "docs":
			runDocs(agentMode)
		case "help", "-h", "--help":
			if agentMode {
				fmt.Print(docs.GetDocs(true))
			} else {
				printHelp()
			}
		default:
			fmt.Printf("Unknown command: %s\n", os.Args[1])
			printHelp()
			os.Exit(1)
		}
	} else {
		runInteractive()
	}
}

func hasFlag(flag string) bool {
	for _, arg := range os.Args {
		if arg == flag {
			return true
		}
	}
	return false
}

func runInteractive() {
	p := tea.NewProgram(ui.NewMainMenuModel(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}

func runBrowse(agentMode bool) {
	if agentMode {
		ui.RunBrowseAgent()
	} else {
		p := tea.NewProgram(ui.NewBrowseModel(), tea.WithAltScreen())
		if _, err := p.Run(); err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
	}
}

func runAdd(agentMode bool) {
	if agentMode {
		ui.RunAddAgent(os.Args[2:])
	} else {
		p := tea.NewProgram(ui.NewAddModel(), tea.WithAltScreen())
		if _, err := p.Run(); err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
	}
}

func runScrape(agentMode bool) {
	if agentMode {
		ui.RunScrapeAgent(os.Args[2:])
	} else {
		p := tea.NewProgram(ui.NewScrapeModel(), tea.WithAltScreen())
		if _, err := p.Run(); err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
	}
}

func runSearch() {
	// Use the new V2 search with grep-like functionality
	ui.RunSearchCLIV2(os.Args[2:])
}

func runGlob() {
	// Run glob command to find files
	ui.GlobCommand(os.Args[2:])
}

func runDocs(agentMode bool) {
	// Show documentation
	fmt.Print(docs.GetDocs(agentMode))
}

func printHelp() {
	fmt.Println(`Various Docs - Documentation Manager with AI-Optimized Search

Usage:
  vd                          Launch interactive mode
  vd browse                   Browse documentation collections
  vd add                      Add a new documentation source
  vd scrape                   Scrape pending documentation
  
Search Commands (grep-like):
  vd search [options] <pattern>              Search all docs
  vd search [options] <project> <pattern>    Search specific project
  vd grep                                    Alias for search
  
  Options:
    -i              Case insensitive (default: true)
    -e              Use regex pattern
    -C <n>          Show n lines of context
    -B <n>          Show n lines before match
    -A <n>          Show n lines after match
    --mode=<mode>   Output mode: content, files_with_matches, count
    --format=json   Output as JSON for AI agents
    --glob="*.md"   File pattern to search
  
File Finding (glob-like):
  vd glob <pattern> [path]    Find files by pattern
  vd find                     Alias for glob
  
Examples:
  vd search "modal gpu"                      # Search for "modal gpu"
  vd search -e "gpu.*function" modal         # Regex search in modal
  vd search --mode=files_with_matches "api"  # Just show files with matches
  vd search -C 5 "error"                      # Show 5 lines of context
  vd search --format=json "query"            # JSON output for AI agents
  vd glob "**/*.md" modal                    # Find all .md files in modal
  
Environment Variables:
  VD_FORMAT=json    Output as JSON (for AI agents like Claude Code)

Description:
  VD provides grep-like and glob-like search capabilities optimized for
  AI assistants like Claude Code, Cursor, and Copilot to efficiently
  search and retrieve documentation.`)
}
