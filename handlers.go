package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
)

func viewPageHandler(w http.ResponseWriter, r *http.Request) {
	// Get the file path from environment variable or use default
	filePath := os.Getenv("FILE_PATH")
	if filePath == "" {
		filePath = "page/the_page.md"
	}

	// Read the markdown file content
	content, err := os.ReadFile(filePath)
	if err != nil {
		// If file doesn't exist, show empty page
		content = []byte("# Welcome to Monopage\n\nThis is a radically minimalistic CMS. Edit this page to add your content.")
	}

	// Convert markdown to HTML
	renderedContent, err := renderMarkdown(string(content))
	if err != nil {
		http.Error(w, "Failed to render content", http.StatusInternalServerError)
		log.Printf("Failed to render markdown: %v", err)
		return
	}

	// Create data structure for template
	data := struct {
		Title           string
		RenderedContent template.HTML
	}{
		Title:           "Monopage",
		RenderedContent: template.HTML(renderedContent),
	}

	// Parse and execute view template
	tmpl, err := template.ParseFiles("templates/view.html")
	if err != nil {
		http.Error(w, "Template error", http.StatusInternalServerError)
		log.Printf("Template error: %v", err)
		return
	}

	tmpl.Execute(w, data)
}

func editPageHandler(w http.ResponseWriter, r *http.Request) {
	// Get the file path from environment variable or use default
	filePath := os.Getenv("FILE_PATH")
	if filePath == "" {
		filePath = "page/the_page.md"
	}

	switch r.Method {
	case "GET":
		// Read the current content of the markdown file
		content, err := os.ReadFile(filePath)
		if err != nil {
			// If file doesn't exist, start with empty content
			content = []byte("")
		}

		// Create data structure for template
		data := struct {
			Content string
		}{
			Content: string(content),
		}

		// Parse and execute edit template
		tmpl, err := template.ParseFiles("templates/edit.html")
		if err != nil {
			http.Error(w, "Template error", http.StatusInternalServerError)
			log.Printf("Template error: %v", err)
			return
		}

		tmpl.Execute(w, data)

	case "POST":
		// Parse form data
		if err := r.ParseForm(); err != nil {
			http.Error(w, "Failed to parse form", http.StatusBadRequest)
			return
		}

		// Get content from form
		content := strings.TrimSpace(r.FormValue("content"))

		// Ensure directory exists
		dir := filePath[:strings.LastIndex(filePath, "/")]
		if err := os.MkdirAll(dir, 0755); err != nil {
			http.Error(w, "Failed to create directory", http.StatusInternalServerError)
			log.Printf("Failed to create directory: %v", err)
			return
		}

		// Write content to file
		if err := os.WriteFile(filePath, []byte(content), 0644); err != nil {
			http.Error(w, "Failed to save page", http.StatusInternalServerError)
			log.Printf("Failed to save page: %v", err)
			return
		}

		// Redirect to view page
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}
