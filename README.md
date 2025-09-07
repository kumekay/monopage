# Monopage

A radically minimalistic CMS implemented in Go that only allows editing and rendering of a single page.
No themes, image uploads, embeddings or anything.
But well, still can be useful somehow.

I personally use it for the "Homepage" with collection URLs and some important notes.

> [!WARNING]
> This application has no built-in authentication. You must bring your own authentication solution to secure the `/edit` endpoint in production environments.

## Features

- Simple markdown editing interface
- CommonMark compliant markdown rendering
- XSS protection through content sanitization
- Docker support for easy deployment
- Includes an example with Traefik integration for Let's Encrypt SSL certificates and OAuth2 Proxy for authentication for the editor endpoint.

## Quick Start

### Using Docker Compose (Recommended)

1. Create a `.env` file with your configuration:
```bash
APP_DOMAIN=your-domain.com
ACME_EMAIL=your-email@domain.com # For TLS certificate
GITHUB_CLIENT_ID=your-github-client-id
GITHUB_CLIENT_SECRET=your-github-client-secret
COOKIE_SECRET=your-random-cookie-secret
GITHUB_USER=your-github-username
EMAIL_DOMAIN=* # or specify a domain like "example.com"
```

2. Run the application:
```bash
docker-compose up -d
```

The application will be available at https://your-domain.com

### Building from Source

1. Clone the repository:
```bash
git clone https://github.com/your-username/monopage.git
cd monopage
```

2. Build the binary:
```bash
go build -o monopage .
```

3. Run the application:
```bash
./monopage
```

By default, the server starts on port 8080.

### URLs

- `/` - View the rendered page
- `/edit` - Edit the page content (protected by GitHub OAuth2)

## Security

Monopage is designed to be run behind a reverse proxy for authentication. The `/edit` endpoint is protected with GitHub OAuth2 when using the Docker Compose setup.

If you need to add authentication to the `/edit` endpoint when running without Docker, you can use a reverse proxy like nginx or Traefik with OAuth2 proxy.

A good option for OAuth2 is https://github.com/oauth2-proxy/oauth2-proxy
