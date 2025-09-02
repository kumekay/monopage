# AGENTS.md

This file provides guidance to AI agents when working with code in this repository.

## Project Overview

Monopage is radically minimalistic CMS implemented in Go. It only allows to edit and render 1 single page.
By default the page is looked in `page/the_page.md` relative to the binary.

The editor is a simple textarea with syntax highlighting for markdown. 
The page is rendered using a CommonMark compliant markdown renderer.
The rendered page is sanitized to prevent XSS attacks.
The project is designed to be run behind a reverse proxy like nginx or traefik, to add authentication for the `edit` endpoint.

### Architecture

- **main.go**: HTTP server setup with Gorilla Mux routing
- **handlers.go**: HTTP handlers for view and edit endpoints
- **templates/**: HTML templates for view and edit pages
- **static/**: Directory for static assets (CSS, JS, images)
- **page/**: Directory for the markdown page content

### Key Dependencies

- `github.com/gorilla/mux` - HTTP routing
- `github.com/yuin/goldmark` - Markdown rendering (CommonMark compliant)

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

- `HOST`: Server host (default: `0.0.0.0`)
- `PORT`: Server port (default: 8080)
- `FILE_PATH`: Path to the markdown file (default: `page/the_page.md`)

### URL Structure

- `/` - Homepage (the only page)
- `/edit` - Editor page

## Security and authentication

If you need to add some authentication to the `/edit` endpoint, then just bring your own reverse proxy.
A good option for oauth https://github.com/oauth2-proxy/oauth2-proxy
