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

	currentURL, err := url.Parse(rawCurrentURL)
	if err != nil {
		fmt.Printf("Error: crawlpage - couldn't parse URL: %s : %v\n", rawCurrentURL, err)
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

	// if _, ok := cfg.pages[normalURL]; ok {
	// 	cfg.pages[normalURL]++
	// 	return
	// }
	// cfg.pages[normalURL] = 1

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
