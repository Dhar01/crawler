package main

import (
	"fmt"
	"log"
	"net/url"
)

func crawlPage(rawBaseURL, rawCurrentURL string, pages map[string]int) {
	baseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		log.Println("Error parsing baseURL:", err)
		return
	}

	currentURL, err := url.Parse(rawCurrentURL)
	if err != nil {
		log.Println("Error parsing currentURL:", err)
		return
	}

	// only crawl pages on the same domain
	if baseURL.Host != currentURL.Host {
		return
	}

	normalURL, err := normalizeURL(rawCurrentURL)
	if err != nil {
		log.Println("Error normalizing URL:", err)
		return
	}

	if _, ok := pages[normalURL]; ok {
		pages[normalURL]++
		return
	}

	pages[normalURL] = 1

	fmt.Println("Crawling:", normalURL)

	html, err := getHTML(rawCurrentURL)
	if err != nil {
		log.Println("Error fetching URL:", rawCurrentURL)
		return
	}

	urls, err := getURLsFromHTML(html, rawBaseURL)
	if err != nil {
		log.Println("Error parsing URLs:", err)
		return
	}

	for _, url := range urls {
		crawlPage(rawBaseURL, url, pages)
	}
}
