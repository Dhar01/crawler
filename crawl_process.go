package main

import (
	"fmt"
	"log"
	"net/url"
)

func (cfg *config) crawlPage(rawCurrentURL string) {
	cfg.concurrencyControl <- struct{}{}
	defer func() {
		<-cfg.concurrencyControl
		cfg.wg.Done()
	}()

	cfg.mu.Lock()

	if len(cfg.pages) >= cfg.maxPages {
		cfg.mu.Unlock()
		return
	}

	cfg.mu.Unlock()

	currentURL, err := url.Parse(rawCurrentURL)
	if err != nil {
		log.Printf("error: crawlPage - couldn't parse URL: %s : %v\n", rawCurrentURL, err)
		return
	}

	// only crawl pages on the same domain
	if cfg.baseURL.Hostname() != currentURL.Hostname() {
		return
	}

	normalURL, err := normalizeURL(rawCurrentURL)
	if err != nil {
		log.Println("Error normalizing URL:", err)
		return
	}

	isFirst := cfg.addPageVisit(normalURL)
	if !isFirst {
		return
	}

	fmt.Println("Crawling:", normalURL)

	html, err := getHTML(rawCurrentURL)
	if err != nil {
		log.Printf("Error: getHTML - fetching URL: %s , %v", rawCurrentURL, err)
		return
	}

	urls, err := getURLsFromHTML(html, cfg.baseURL)
	if err != nil {
		log.Println("Error parsing URLs:", err)
		return
	}

	for _, url := range urls {
		cfg.wg.Add(1)
		go cfg.crawlPage(url)
	}
}

func (cfg *config) addPageVisit(normalizeURL string) (isFirst bool) {
	cfg.mu.Lock()
	defer cfg.mu.Unlock()

	if _, visited := cfg.pages[normalizeURL]; visited {
		cfg.pages[normalizeURL]++
		return false
	}

	cfg.pages[normalizeURL] = 1
	return true
}
