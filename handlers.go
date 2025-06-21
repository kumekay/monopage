package main

import (
	"html/template"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

func createPageHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		tmpl, err := template.ParseFiles("templates/create.html")
		if err != nil {
			http.Error(w, "Template error", http.StatusInternalServerError)
			log.Printf("Template error: %v", err)
			return
		}
		tmpl.Execute(w, nil)
		
	case "POST":
		if err := r.ParseForm(); err != nil {
			http.Error(w, "Failed to parse form", http.StatusBadRequest)
			return
		}
		
		title := strings.TrimSpace(r.FormValue("title"))
		content := strings.TrimSpace(r.FormValue("content"))
		
		if title == "" || content == "" {
			http.Error(w, "Title and content are required", http.StatusBadRequest)
			return
		}
		
		slug := generateSlug(title)
		editToken := generateEditToken()
		
		page, err := createPage(title, content, slug, editToken)
		if err != nil {
			http.Error(w, "Failed to create page", http.StatusInternalServerError)
			log.Printf("Failed to create page: %v", err)
			return
		}
		
		http.SetCookie(w, &http.Cookie{
			Name:     "edit_token_" + slug,
			Value:    editToken,
			Path:     "/",
			HttpOnly: true,
			MaxAge:   86400 * 365,
		})
		
		http.Redirect(w, r, "/"+page.Slug, http.StatusSeeOther)
	}
}

func viewPageHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	slug := vars["slug"]
	
	page, err := getPageBySlug(slug)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	
	renderedContent, err := renderMarkdown(page.Content)
	if err != nil {
		http.Error(w, "Failed to render content", http.StatusInternalServerError)
		log.Printf("Failed to render markdown: %v", err)
		return
	}
	
	data := struct {
		Page            *Page
		RenderedContent template.HTML
		CanEdit         bool
	}{
		Page:            page,
		RenderedContent: template.HTML(renderedContent),
		CanEdit:         canEditPage(r, page),
	}
	
	tmpl, err := template.ParseFiles("templates/view.html")
	if err != nil {
		http.Error(w, "Template error", http.StatusInternalServerError)
		log.Printf("Template error: %v", err)
		return
	}
	
	tmpl.Execute(w, data)
}

func editPageHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fullSlug := vars["slug"]
	slug := strings.TrimSuffix(fullSlug, "_edit")
	
	page, err := getPageBySlug(slug)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	
	switch r.Method {
	case "GET":
		data := struct {
			Page *Page
		}{
			Page: page,
		}
		
		tmpl, err := template.ParseFiles("templates/edit.html")
		if err != nil {
			http.Error(w, "Template error", http.StatusInternalServerError)
			log.Printf("Template error: %v", err)
			return
		}
		
		tmpl.Execute(w, data)
		
	case "POST":
		if err := r.ParseForm(); err != nil {
			http.Error(w, "Failed to parse form", http.StatusBadRequest)
			return
		}
		
		title := strings.TrimSpace(r.FormValue("title"))
		content := strings.TrimSpace(r.FormValue("content"))
		
		if title == "" || content == "" {
			http.Error(w, "Title and content are required", http.StatusBadRequest)
			return
		}
		
		if err := updatePage(slug, title, content); err != nil {
			http.Error(w, "Failed to update page", http.StatusInternalServerError)
			log.Printf("Failed to update page: %v", err)
			return
		}
		
		http.Redirect(w, r, "/"+slug, http.StatusSeeOther)
	}
}

func canEditPage(r *http.Request, page *Page) bool {
	cookie, err := r.Cookie("edit_token_" + page.Slug)
	if err != nil {
		return false
	}
	return cookie.Value == page.EditToken
}