package main

import (
	"net/url"
	"strings"
)

func normalizeURL(address string) (string, error) {
	urlStr, err := url.Parse(address)
    if err != nil {
        return "", err
    }

    trimmedPath := strings.TrimSuffix(urlStr.Path, "/")

    normStr := urlStr.Host + trimmedPath
    return normStr,nil
}
