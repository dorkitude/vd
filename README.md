# VD - Various Docs 📚

A local repository system for storing and browsing developer documentation in Markdown format, making it accessible to local AI agents like Claude Code.

## Overview

VD (Various Docs) is a documentation aggregator that creates local copies of various online documentation sources. It provides an interactive CLI built with [Bubble Tea](https://github.com/charmbracelet/bubbletea) for browsing and managing these documentation collections.

## Features

- 📚 **Browse Documentation**: Navigate through documentation with beautiful terminal UI
- ➕ **Add Sources**: Add new documentation sources (Mintlify sites, GitHub repos)
- 🕷️ **Smart Scraping**: Automatically scrape and convert documentation to Markdown
- 🎨 **Beautiful Rendering**: View markdown files with Glamour styling
- 🤖 **AI-Friendly**: Optimized for local AI assistants to read and reference

## Installation

```bash
# Clone the repository
git clone https://github.com/dorkitude/vd.git
cd vd

# Build and install
make dev

# Or build only
make build
```

## Usage

```bash
# Launch interactive mode
vd

# Browse documentation collections
vd browse

# Add a new documentation source
vd add

# Scrape pending documentation
vd scrape

# Get help
vd help
```

## Interactive Navigation

- **Main Menu**: Choose between browsing, adding, or scraping
- **Browse Mode**: 
  - Select a project → View details → Browse files → Read with Glamour
  - Use `/` to search, `Enter` to select, `Esc` to go back
- **Add Mode**: Input folder name, title, description, and source URL
- **Scrape Mode**: Select pending projects and scrape their documentation

## Structure

```
vd/
├── content/                 # All documentation collections
│   ├── [project-name]/     # Individual documentation project
│   │   ├── vd.json        # Metadata for this collection
│   │   └── ...            # Documentation files
├── cmd/vd/                 # CLI entry point
├── internal/               # Core application logic
│   ├── models/            # Data structures
│   ├── ui/                # Bubble Tea UI components
│   └── scraper/           # Scraping logic
└── scripts/               # Helper scripts (Python scraper)
```

## Metadata Format (vd.json)

Each documentation collection contains a `vd.json` file with:
- **title**: Documentation title
- **description**: What this documentation covers
- **root_url**: The base URL that was scraped
- **files**: Array of file metadata with paths, URLs, and timestamps
- **metadata**: Additional metadata like doc_type and status

## Supported Documentation Formats

- **Mintlify Sites**: Modern documentation platforms
- **GitHub Repositories**: README files and documentation folders
- More formats can be added by extending the scraper

## Development

Built with:
- **Go** - Core application
- **[Bubble Tea](https://github.com/charmbracelet/bubbletea)** - Terminal UI framework
- **[Glamour](https://github.com/charmbracelet/glamour)** - Beautiful markdown rendering
- **[Lipgloss](https://github.com/charmbracelet/lipgloss)** - Terminal styling
- **Python** - Web scraping scripts

## Contributing

Feel free to add support for additional documentation formats or improve the CLI experience!

## License

MIT