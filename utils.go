package main

import (
	"bytes"
	"crypto/rand"
	"encoding/hex"
	"regexp"
	"strings"

	"github.com/yuin/goldmark"
)

func generateSlug(title string) string {
	slug := strings.ToLower(title)

	reg := regexp.MustCompile(`[^a-z0-9\s-]`)
	slug = reg.ReplaceAllString(slug, "")

	slug = strings.ReplaceAll(slug, " ", "_")

	reg = regexp.MustCompile(`_+`)
	slug = reg.ReplaceAllString(slug, "_")

	slug = strings.Trim(slug, "_")

	if slug == "" {
		slug = "untitled"
	}

	randomHex := generateRandomHex(6)
	slug = slug + "_" + randomHex

	for pageExists(slug) {
		randomHex = generateRandomHex(6)
		lastUnderscore := strings.LastIndex(slug, "_")
		if lastUnderscore != -1 {
			slug = slug[:lastUnderscore] + "_" + randomHex
		} else {
			slug = slug + "_" + randomHex
		}
	}

	return slug
}

func generateRandomHex(length int) string {
	bytes := make([]byte, length/2)
	if _, err := rand.Read(bytes); err != nil {
		return "000000"
	}
	return hex.EncodeToString(bytes)[:length]
}

func generateEditToken() string {
	return generateRandomHex(32)
}

func renderMarkdown(content string) (string, error) {
	var buf bytes.Buffer
	if err := goldmark.Convert([]byte(content), &buf); err != nil {
		return "", err
	}
	return buf.String(), nil
}
