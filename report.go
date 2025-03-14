package main

import (
	"fmt"
	"sort"
)

func printReport(pagesMap map[string]int, baseURL string) {
	fmt.Println("===================================")
	fmt.Printf("REPORT for %s\n", baseURL)
	fmt.Println("===================================")

	var pages []PageCount

	for url, count := range pagesMap {
		pages = append(pages, PageCount{
			URL:   url,
			Count: count,
		})
	}

	sortResult(pages)

	for _, page := range pages {
		fmt.Printf("Found %d internal links to %s\n", page.Count, page.URL)
	}
}

type PageCount struct {
	URL   string
	Count int
}

func sortResult(pages []PageCount) {
	sort.Slice(pages, func(i, j int) bool {
		// if counts differ, sort by count
		if pages[i].Count != pages[j].Count {
			return pages[i].Count > pages[j].Count
		}
		// if counts are equal, sorting alphabetically by URL
		return pages[i].URL < pages[j].URL
	})
}
