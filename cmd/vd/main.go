package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/dorkitude/vd/internal/ui"
)

func main() {
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "browse":
			runBrowse()
		case "add":
			runAdd()
		case "scrape":
			runScrape()
		case "help", "-h", "--help":
			printHelp()
		default:
			fmt.Printf("Unknown command: %s\n", os.Args[1])
			printHelp()
			os.Exit(1)
		}
	} else {
		runInteractive()
	}
}

func runInteractive() {
	p := tea.NewProgram(ui.NewMainMenuModel(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}

func runBrowse() {
	p := tea.NewProgram(ui.NewBrowseModel(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}

func runAdd() {
	p := tea.NewProgram(ui.NewAddModel(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}

func runScrape() {
	p := tea.NewProgram(ui.NewScrapeModel(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}

func printHelp() {
	fmt.Println(`Various Docs

Usage:
  vd              Launch interactive mode
  vd browse       Browse documentation collections
  vd add          Add a new documentation source
  vd scrape       Scrape pending documentation
  vd help         Show this help message

Description:
  VD is a local documentation manager that stores and organizes
  developer documentation in Markdown format for easy access by
  local AI agents and developers.

Examples:
  vd                # Launch the main menu
  vd browse         # Jump directly to browse mode
  vd add            # Jump directly to add a new project`)
}
