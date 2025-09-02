package main

import (
	"bytes"

	"github.com/yuin/goldmark"
)

func renderMarkdown(content string) (string, error) {
	var buf bytes.Buffer
	if err := goldmark.Convert([]byte(content), &buf); err != nil {
		return "", err
	}
	return buf.String(), nil
}
