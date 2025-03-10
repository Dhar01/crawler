package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func getHTML(rawURL string) (string, error) {
	resp, err := http.Get(rawURL)
	if err != nil {
		return "", fmt.Errorf("got network error: %v", err)
	}
	defer resp.Body.Close()

	// if status code is 400+
	if resp.StatusCode > 399 {
		return "", fmt.Errorf("got HTTP error: %s", resp.Status)
	}

	// if the header doesn't contain text/html, return error
	contentType := resp.Header.Get("Content-Type")
	if !strings.Contains(contentType, "text/html") {
		return "", fmt.Errorf("content-type is not text/html")
	}

	content, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(content), nil
}
