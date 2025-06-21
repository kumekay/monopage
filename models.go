package main

import (
	"database/sql"
	"time"
)

type Page struct {
	ID        int       `json:"id"`
	Slug      string    `json:"slug"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	EditToken string    `json:"-"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func createPage(title, content, slug, editToken string) (*Page, error) {
	query := `
		INSERT INTO pages (title, content, slug, edit_token)
		VALUES (?, ?, ?, ?)
	`
	result, err := db.Exec(query, title, content, slug, editToken)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	return getPageByID(int(id))
}

func getPageBySlug(slug string) (*Page, error) {
	query := `
		SELECT id, slug, title, content, edit_token, created_at, updated_at
		FROM pages
		WHERE slug = ?
	`
	
	page := &Page{}
	err := db.QueryRow(query, slug).Scan(
		&page.ID, &page.Slug, &page.Title, &page.Content,
		&page.EditToken, &page.CreatedAt, &page.UpdatedAt,
	)
	
	if err != nil {
		return nil, err
	}
	
	return page, nil
}

func getPageByID(id int) (*Page, error) {
	query := `
		SELECT id, slug, title, content, edit_token, created_at, updated_at
		FROM pages
		WHERE id = ?
	`
	
	page := &Page{}
	err := db.QueryRow(query, id).Scan(
		&page.ID, &page.Slug, &page.Title, &page.Content,
		&page.EditToken, &page.CreatedAt, &page.UpdatedAt,
	)
	
	if err != nil {
		return nil, err
	}
	
	return page, nil
}

func updatePage(slug, title, content string) error {
	query := `
		UPDATE pages
		SET title = ?, content = ?, updated_at = CURRENT_TIMESTAMP
		WHERE slug = ?
	`
	_, err := db.Exec(query, title, content, slug)
	return err
}

func pageExists(slug string) bool {
	query := `SELECT 1 FROM pages WHERE slug = ? LIMIT 1`
	var exists int
	err := db.QueryRow(query, slug).Scan(&exists)
	return err != sql.ErrNoRows
}