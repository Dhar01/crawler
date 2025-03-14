package main

import (
	"fmt"
	"log"
	"os"

	"strconv"
)

func main() {
	args := os.Args

	if len(args) < 4 {
		fmt.Println("no website provided")
		os.Exit(1)
	} else if len(args) > 4 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	} else {
		fmt.Printf("starting crawl of: %s\n", args[1])
	}

	website := args[1]
	maxConcur := args[2]
	pages := args[3]

	maxConcurrency, err := strconv.Atoi(maxConcur)
	if err != nil {
		log.Println("maxConcurrency isn't integer")
		os.Exit(1)
	}

	maxPages, err := strconv.Atoi(pages)
	if err != nil {
		log.Println("maxPages isn't integer")
		os.Exit(1)
	}

	cfg, err := configure(website, maxConcurrency, maxPages)
	if err != nil {
		fmt.Printf("Error - configure: %v", err)
		return
	}

	fmt.Println("Starting crawl of:", website)

	cfg.wg.Add(1)
	go cfg.crawlPage(website)
	cfg.wg.Wait()

	printReport(cfg.pages, website)
}
