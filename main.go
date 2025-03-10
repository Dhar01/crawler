package main

import (
	"fmt"
	"log"
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

	result, err := getHTML(website)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(result)
}
