package docs

import (
	"fmt"
	"strings"
)

// GetDocs returns documentation in either human or agent format
func GetDocs(agentMode bool) string {
	if agentMode {
		return getAgentDocs()
	}
	return getHumanDocs()
}

// getAgentDocs returns compressed, agent-friendly documentation
func getAgentDocs() string {
	return `VD Documentation for AI Agents

COMMANDS:
search <pattern> [project] - Full-text search. Flags: --format=json --mode={content|files_with_matches|count} -C <n> -e (regex)
glob <pattern> [path] - Find files by pattern. Returns paths sorted by mod time.
browse --agent - List all projects and files as JSON
scrape --agent [project] - Scrape docs. Returns status JSON.
add --agent <url> <name> - Add doc source. Non-interactive.

SEARCH EXAMPLES:
vd search --format=json "modal gpu" - JSON results with context
vd search --mode=files_with_matches "api" - Just matching files
vd search -e "func.*gpu" modal - Regex in specific project
vd glob "**/*.md" - Find all markdown files

OUTPUT FORMATS:
- Add --format=json or VD_FORMAT=json for JSON output
- All commands support --agent for terse, parseable output
- Search returns: {matches:[{file,project,line,column,match,context}]}
- Browse returns: {projects:[{name,path,files:[]}]}

PATHS:
Content stored in: content/<project>/*.md
Metadata in: content/<project>/vd.json

USAGE:
Best for: Finding specific documentation, listing available docs, checking scrape status.
Chain commands: First glob/browse to find files, then search for specific content.`
}

// getHumanDocs returns human-friendly documentation
func getHumanDocs() string {
	return `
╔════════════════════════════════════════════════════════════════════╗
║                     VD - Various Docs                              ║
║           Your AI-Powered Documentation Manager                    ║
╚════════════════════════════════════════════════════════════════════╝

📚 OVERVIEW
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
VD stores and organizes developer documentation locally, optimized for
AI assistants like Claude Code, Cursor, and GitHub Copilot.

🎯 QUICK START
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
  vd                    Launch interactive menu
  vd search "query"     Search all documentation
  vd browse            Browse documentation library
  vd scrape            Scrape pending documentation

🔍 SEARCH COMMANDS
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
Search with grep-like functionality:

  vd search "pattern"                 Search all projects
  vd search project "pattern"         Search specific project
  vd search -e "regex.*pattern"       Use regex
  vd search -C 5 "pattern"           Show 5 lines of context
  
Output Modes:
  --mode=content              Show matches with context (default)
  --mode=files_with_matches   Just list matching files
  --mode=count               Show match counts per file
  
Context Options:
  -C <n>    Lines before and after match
  -B <n>    Lines before match
  -A <n>    Lines after match
  
Other Options:
  -i        Case insensitive (default: true)
  -e        Enable regex mode
  --glob    File pattern (default: "*.md")
  --format=json   JSON output for AI agents

📁 FILE FINDING
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
Find files with glob patterns:

  vd glob "*.md"              Find all markdown files
  vd glob "**/*.py" project   Find Python files in project
  vd find "test_*"           Alias for glob

🕷️ DOCUMENTATION MANAGEMENT
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
  vd add              Add new documentation source (interactive)
  vd scrape          Scrape pending documentation
  vd browse          Browse and read documentation

🤖 AI AGENT MODE
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
All commands support --agent flag for terse, machine-readable output:

  vd docs --agent                    Compressed documentation
  vd browse --agent                  JSON project listing
  vd scrape --agent                  JSON scrape status
  vd search --format=json "query"    JSON search results

Environment Variables:
  VD_FORMAT=json      Default to JSON output
  VD_AGENT=true      Default to agent mode

💡 EXAMPLES
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
  # Search Modal docs for GPU references
  vd search modal "gpu"
  
  # Find all examples in Modal docs
  vd glob "**/examples/*.md" modal
  
  # Get JSON results for AI processing
  vd search --format=json "api endpoint"
  
  # Show just files containing "webhook"
  vd search --mode=files_with_matches "webhook"
  
  # Search with regex and context
  vd search -e "class.*Model" -C 3

📂 PROJECT STRUCTURE
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
  content/
  ├── modal/                 # Modal documentation
  │   ├── vd.json           # Project metadata
  │   └── *.md              # Documentation files
  └── [project]/            # Other documentation

⌨️ KEYBOARD SHORTCUTS
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
  ↑/↓, j/k     Navigate items
  Enter        Select/Open
  /            Search mode
  Esc, q       Go back/Quit
  ?            Show help
  Tab          Switch panes

🔗 MORE INFORMATION
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
  GitHub:  https://github.com/dorkitude/vd
  Author:  Kyle Wild (Chef)
`
}

// GetCommandHelp returns help for a specific command in agent or human format
func GetCommandHelp(command string, agentMode bool) string {
	if agentMode {
		return getAgentCommandHelp(command)
	}
	return getHumanCommandHelp(command)
}

func getAgentCommandHelp(command string) string {
	helps := map[string]string{
		"search": `search <pattern> [project] [flags]
Flags: --format=json --mode={content|files_with_matches|count} -C <n> -B <n> -A <n> -e -i --glob="*.md"
Output: JSON with matches array containing file,project,line,column,match,context`,
		
		"glob": `glob <pattern> [path]
Find files matching pattern. Returns paths sorted by modification time.
Output: File paths, one per line. With --format=json: array of {path,project,size,mod_time}`,
		
		"browse": `browse --agent
List all projects and files.
Output: JSON {projects:[{name,path,file_count,total_size}]}`,
		
		"scrape": `scrape --agent [project]
Scrape documentation. Non-interactive in agent mode.
Output: JSON {status,projects_scraped,errors}`,
		
		"add": `add --agent <url> <name> [--type=mintlify|modal]
Add documentation source. Non-interactive in agent mode.
Output: JSON {success,project,message}`,
	}
	
	if help, ok := helps[command]; ok {
		return help
	}
	return fmt.Sprintf("Unknown command: %s", command)
}

func getHumanCommandHelp(command string) string {
	helps := map[string]string{
		"search": formatCommandHelp("SEARCH", `
Search through documentation with grep-like functionality.

USAGE:
  vd search [options] <pattern>
  vd search [options] <project> <pattern>

OPTIONS:
  -i              Case insensitive (default: true)
  -e              Use regex pattern
  -C <n>          Show n lines of context
  -B <n>          Show n lines before match
  -A <n>          Show n lines after match
  --mode=<mode>   Output mode: content, files_with_matches, count
  --format=json   Output as JSON
  --glob="*.md"   File pattern to search

EXAMPLES:
  vd search "modal gpu"
  vd search -e "func.*gpu" modal
  vd search --mode=files_with_matches "api"
  vd search -C 5 "error"`),
		
		"glob": formatCommandHelp("GLOB", `
Find files matching a glob pattern.

USAGE:
  vd glob <pattern> [path]
  vd find <pattern> [path]  (alias)

PATTERNS:
  *.md           Match .md files in current dir
  **/*.md        Match .md files recursively
  test_*.js      Match test files

EXAMPLES:
  vd glob "*.md" modal
  vd glob "**/*.py"
  vd find "examples/*.md"`),
		
		"browse": formatCommandHelp("BROWSE", `
Browse documentation collections interactively.

USAGE:
  vd browse         Interactive mode
  vd browse --agent JSON listing mode

NAVIGATION:
  ↑/↓ or j/k    Navigate items
  Enter         Open item
  /             Search
  q or Esc      Go back`),
	}
	
	if help, ok := helps[command]; ok {
		return help
	}
	return fmt.Sprintf("\n❌ Unknown command: %s\n\nUse 'vd help' to see available commands.\n", command)
}

func formatCommandHelp(title, content string) string {
	border := strings.Repeat("─", 60)
	return fmt.Sprintf(`
╔════════════════════════════════════════════════════════════╗
║  %s%-58s  ║
╚════════════════════════════════════════════════════════════╝
%s
%s
`, "", title, content, border)
}