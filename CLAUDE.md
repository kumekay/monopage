# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

Monopage is a minimalistic CMS for markdown files implemented in Go. It's similar to Telegraph in its simplicity but with a slightly different workflow.

### Core Workflow
1. User opens homepage and enters title and markdown body
2. Clicks "Publish" to create page with auto-generated URL (title words + random hex)
3. Page displays rendered markdown with edit link (if user has edit cookie)
4. Edit page uses separate hex token + `_edit` suffix for access control

### Architecture

- **main.go**: HTTP server setup with Gorilla Mux routing
- **database.go**: SQLite database initialization and connection
- **models.go**: Page struct and CRUD operations  
- **handlers.go**: HTTP handlers for create/view/edit endpoints
- **utils.go**: URL generation, random hex tokens, Goldmark markdown rendering
- **templates/**: HTML templates for create, view, and edit pages
- **static/**: Directory for static assets (CSS, JS, images)

### Database Schema
```sql
pages (
  id INTEGER PRIMARY KEY,
  slug TEXT UNIQUE,
  title TEXT,
  content TEXT,
  edit_token TEXT,
  created_at DATETIME,
  updated_at DATETIME
)
```

### Key Dependencies
- `github.com/gorilla/mux` - HTTP routing
- `github.com/yuin/goldmark` - Markdown rendering (CommonMark compliant)
- `modernc.org/sqlite` - Pure Go SQLite driver

## Development Commands

### Setup and Dependencies
```bash
go mod download
```

### Run Development Server
```bash
go run .
# Server starts on :8080 (or PORT env var)
```

### Build
```bash
go build -o monopage .
```

### Run Built Binary
```bash
./monopage
```

### Docker
```bash
# Build image
docker build -t monopage .

# Run container
docker run -p 8080:8080 monopage
```

## Configuration

### Environment Variables
- `PORT`: Server port (default: 8080)
- `DB_PATH`: SQLite database path (default: ./monopage.db)

### URL Structure
- `/` - Homepage (create new page)
- `/{slug}` - View page  
- `/{slug}_edit` - Edit page (requires edit cookie)

## Security Model

- Edit permissions managed via HTTP-only cookies
- Edit tokens are 32-character random hex strings
- Each page has unique edit token stored in database
- Edit URLs include separate random hex suffix for obscurity