package main

import (
	"fmt"
	"net/url"
	"strings"
)

func normalizeURL(address string) (string, error) {
	urlStr, err := url.Parse(address)
	if err != nil {
		return "", fmt.Errorf("couldn't parse URL: %w", err)
	}

	trimmedPath := strings.TrimSuffix(urlStr.Path, "/")

	normStr := urlStr.Host + trimmedPath
	return strings.ToLower(normStr), nil
}

func getURLsFromHTML(htmlBody, rawBaseURL string) ([]string, error) {
	return nil, nil
}
