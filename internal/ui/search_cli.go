package ui

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/dorkitude/vd/internal/search"
)

// SearchCLIArgs holds parsed command line arguments for search
type SearchCLIArgs struct {
	Pattern      string
	Path         string
	OutputMode   string
	BeforeLines  int
	AfterLines   int
	ContextLines int
	MaxResults   int
	IgnoreCase   bool
	ShowLineNums bool
	UseRegex     bool
	FilePattern  string
	Format       string
}

// RunSearchCLIV2 runs the improved grep-like search from command line
func RunSearchCLIV2(args []string) {
	// Create a new flag set for search command
	searchFlags := flag.NewFlagSet("search", flag.ExitOnError)
	
	// Define flags (similar to grep/ripgrep)
	var opts SearchCLIArgs
	searchFlags.StringVar(&opts.OutputMode, "mode", "content", "Output mode: files_with_matches, content, count")
	searchFlags.IntVar(&opts.BeforeLines, "B", 0, "Lines before match")
	searchFlags.IntVar(&opts.AfterLines, "A", 0, "Lines after match")
	searchFlags.IntVar(&opts.ContextLines, "C", 2, "Lines before and after match")
	searchFlags.IntVar(&opts.MaxResults, "max", 100, "Maximum results")
	searchFlags.BoolVar(&opts.IgnoreCase, "i", true, "Case insensitive search")
	searchFlags.BoolVar(&opts.ShowLineNums, "n", true, "Show line numbers")
	searchFlags.BoolVar(&opts.UseRegex, "e", false, "Use regex pattern")
	searchFlags.StringVar(&opts.FilePattern, "glob", "*.md", "File pattern to search")
	searchFlags.StringVar(&opts.Format, "format", "", "Output format: json or text")
	
	// Custom usage
	searchFlags.Usage = func() {
		fmt.Fprintf(os.Stderr, `Usage: vd search [options] [path] <pattern>

Search through documentation with grep-like functionality.

Examples:
  vd search "modal gpu"                    # Search all docs for "modal gpu"
  vd search modal "gpu"                     # Search modal project for "gpu"
  vd search -e "gpu.*function" modal       # Regex search in modal
  vd search --mode=files_with_matches "api" # Just show matching files
  vd search -C 5 "error"                    # Show 5 lines of context
  vd search --format=json "query"          # JSON output for AI agents

Options:
`)
		searchFlags.PrintDefaults()
	}
	
	// Parse flags
	searchFlags.Parse(args)
	
	// Get positional arguments
	posArgs := searchFlags.Args()
	
	// Handle environment variables (for backward compatibility)
	if os.Getenv("VD_FORMAT") == "json" {
		opts.Format = "json"
	}
	
	// Parse positional arguments
	switch len(posArgs) {
	case 0:
		// Interactive mode - launch TUI
		RunInteractiveSearch()
		return
	case 1:
		// Search all projects
		opts.Pattern = posArgs[0]
		opts.Path = ""
	case 2:
		// Could be: path pattern OR pattern with spaces
		// Try to determine if first arg is a project
		if isProject(posArgs[0]) {
			opts.Path = posArgs[0]
			opts.Pattern = posArgs[1]
		} else {
			// Assume it's a pattern with spaces
			opts.Pattern = strings.Join(posArgs, " ")
			opts.Path = ""
		}
	default:
		// First is path, rest is pattern
		opts.Path = posArgs[0]
		opts.Pattern = strings.Join(posArgs[1:], " ")
	}
	
	// Run search
	engine := search.NewSearchEngineV2("content")
	
	// Convert to search options
	searchOpts := search.GrepOptions{
		Pattern:      opts.Pattern,
		Path:         opts.Path,
		OutputMode:   search.OutputMode(opts.OutputMode),
		BeforeLines:  opts.BeforeLines,
		AfterLines:   opts.AfterLines,
		ContextLines: opts.ContextLines,
		MaxResults:   opts.MaxResults,
		IgnoreCase:   opts.IgnoreCase,
		ShowLineNums: opts.ShowLineNums,
		IsRegex:      opts.UseRegex,
		FilePattern:  opts.FilePattern,
	}
	
	// Perform search
	results, err := engine.Grep(searchOpts)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	
	// Output results
	if opts.Format == "json" {
		data, _ := json.MarshalIndent(results, "", "  ")
		fmt.Println(string(data))
	} else {
		fmt.Print(search.FormatResponse(results, "text"))
	}
}

// RunInteractiveSearch launches the interactive search UI
func RunInteractiveSearch() {
	// This would launch the TUI search interface
	// For now, just show help
	fmt.Println("Interactive search mode not yet implemented in V2.")
	fmt.Println("Use: vd search <pattern> for command-line search")
}

// isProject checks if a string is likely a project name
func isProject(name string) bool {
	// Check if directory exists in content/
	info, err := os.Stat("content/" + name)
	if err == nil && info.IsDir() {
		return true
	}
	
	// Could also check for common patterns that indicate it's not a project
	// (e.g., contains spaces, special regex characters, etc.)
	if strings.Contains(name, " ") || strings.Contains(name, "*") || strings.Contains(name, "[") {
		return false
	}
	
	return false
}

// GlobCommand implements the glob functionality
func GlobCommand(args []string) {
	if len(args) < 1 {
		fmt.Fprintf(os.Stderr, "Usage: vd glob <pattern> [path]\n")
		fmt.Fprintf(os.Stderr, "Example: vd glob '**/*.md' modal\n")
		os.Exit(1)
	}
	
	pattern := args[0]
	path := ""
	if len(args) > 1 {
		path = args[1]
	}
	
	engine := search.NewSearchEngineV2("content")
	results, err := engine.Glob(search.GlobOptions{
		Pattern:    pattern,
		Path:       path,
		SortByTime: true,
	})
	
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	
	// Output results
	if os.Getenv("VD_FORMAT") == "json" {
		data, _ := json.MarshalIndent(results, "", "  ")
		fmt.Println(string(data))
	} else {
		for _, result := range results {
			fmt.Println(result.Path)
		}
	}
}