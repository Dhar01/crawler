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

	pages := make(map[string]int)

	// result, err := getHTML(website)
	// if err != nil {
	// 	log.Println(err)
	// }

	crawlPage(website, website, pages)

	fmt.Println("Results:")
	for url, count := range pages {
		fmt.Printf("%s: %d\n", url, count)
	}
}
