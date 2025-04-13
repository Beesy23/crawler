package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args

	if len(args) < 2 {
		fmt.Print("no website provided\n")
		os.Exit(1)
	}
	if len(args) > 2 {
		fmt.Print("too many arguments provided\n")
		os.Exit(1)
	}

	rawBaseURL := args[1]
	fmt.Printf("starting crawl of: %s\n", rawBaseURL)

	pages := make(map[string]int)

	crawlPage(rawBaseURL, rawBaseURL, pages)

	fmt.Println("\nCrawling results:")
	for url, count := range pages {
		fmt.Printf("%s:%d\n", url, count)
	}
}
