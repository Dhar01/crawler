package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args

	if len(args) < 2 {
		fmt.Println("no website provided")
		os.Exit(1)
	} else if len(args) > 2 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	} else {
		fmt.Printf("starting crawl of: %s\n", args[1])
	}

	website := args[1]

	const maxConcurrency = 3
	cfg, err := configure(website, maxConcurrency)
	if err != nil {
		fmt.Printf("Error - configure: %v", err)
		return
	}

	fmt.Println("Starting crawl of:", website)

	cfg.wg.Add(1)
	go cfg.crawlPage(website)
	cfg.wg.Wait()

	for url, count := range cfg.pages {
		fmt.Printf("%s: %d\n", url, count)
	}
}
