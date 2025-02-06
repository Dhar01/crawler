package main

import (
	"fmt"
	"net/url"
	"strings"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
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
	doc, err := html.Parse(strings.NewReader(htmlBody))
	if err != nil {
		return nil, err
	}

	var elementList []string

	for n := range doc.Descendants() {
		if n.Type == html.ElementNode && n.DataAtom == atom.A {
			for _, a := range n.Attr {
				if a.Key == "href" {
					if !strings.Contains(a.Val, "https") {
						a.Val = rawBaseURL + a.Val
					}

					elementList = append(elementList, a.Val)
				}
			}
		}
	}

	return elementList, nil
}
