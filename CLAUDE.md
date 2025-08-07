# VD Project Notes for Claude Code

This file provides context for Claude Code when working on the VD (Various Docs) project.

## Project Overview

VD is a documentation aggregator with a beautiful TUI, designed to store developer documentation locally for AI assistants. It's built in Go with Bubble Tea for the UI and supports multiple documentation scrapers.

## Development Setup

### Quick Development Workflow
- Run `make dev` to install a wrapper script that always runs the latest code
- This means you can edit Go files and immediately test without recompiling
- The wrapper script is installed at `~/bin/vd` and runs `go run cmd/vd/main.go`

### Building for Production
- Run `make build` to create a compiled binary
- Run `make install` to install the compiled binary to `~/bin/vd`

## Project Architecture

### Core Components
1. **UI Layer** (`internal/ui/`)
   - Bubble Tea-based terminal UI
   - Main menu, browse, add, and scrape interfaces
   
2. **Scrapers** (`internal/scraper/`)
   - `mintlify.go` - Handles Mintlify documentation sites
   - `modal.go` - Modal-specific scraper (custom implementation)
   - Python scrapers in `scripts/` for complex HTML parsing

3. **Models** (`internal/models/`)
   - Data structures for metadata and documentation

## Scraper System

### How Scrapers Work
1. UI detects doc type from URL or metadata
2. Routes to appropriate Go scraper or Python script
3. Scraper fetches pages and converts to Markdown
4. Saves to `content/<project_name>/` with `vd.json` metadata

### Adding New Scrapers
1. Create Go implementation in `internal/scraper/`
2. Add detection logic in `internal/ui/scrape.go`
3. For complex HTML parsing, can add Python script in `scripts/`

## Known Issues & TODOs

### Current Limitations
- Modal scraper creates placeholder content (HTML-to-Markdown conversion needs proper parser)
- Python scrapers have better HTML extraction but slower performance
- Interactive UI doesn't work in non-TTY environments (use for testing only)

### Improvement Opportunities
- Add proper HTML-to-Markdown conversion using a library like `goquery`
- Implement concurrent scraping for faster performance
- Add progress bars for long-running scrapes
- Support for incremental updates (only scrape changed pages)

## Testing

### Manual Testing
```bash
# Test scraper directly (bypass UI)
go run test_scraper.go

# Test specific commands
./vd browse
./vd scrape
./vd add
```

### Important Files
- `cmd/vd/main.go` - Entry point
- `internal/ui/scrape.go` - Scraper selection logic
- `internal/scraper/modal.go` - Modal docs scraper
- `Makefile` - Build automation

## Code Style
- Use Go idioms and conventions
- Keep UI code in `internal/ui/`
- Keep scraper logic in `internal/scraper/`
- Prefer Go implementations over Python when feasible

## Project Folders

### Documentation Storage
- `content/modal/` - Modal docs scraped with Go scraper (placeholder content)
- `content/modal_python/` - Modal docs scraped with Python (full content)
- Each folder contains `vd.json` with metadata and scraped `.md` files

## Notes for Kyle (Chef)
- Prefers `make dev` for quick iteration without compilation
- Uses `~/bin/` for personal binaries (not in PATH by default)
- Likes clean, self-contained commits
- Overmind user for running services (logs in `overmind.log.local`)