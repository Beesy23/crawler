package main

import (
	"fmt"
	"net/url"
	"os"
	"sync"
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
	baseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		fmt.Printf("Error parsing URL: %v\n", err)
		return
	}
	fmt.Printf("starting crawl of: %s\n", rawBaseURL)

	cfg := &config{
		pages:              make(map[string]int),
		baseURL:            baseURL,
		mu:                 &sync.Mutex{},
		concurrencyControl: make(chan struct{}, 10),
		wg:                 &sync.WaitGroup{},
	}

	cfg.crawlPage(rawBaseURL)
	cfg.wg.Wait()

	fmt.Println("\nCrawling results:")
	for url, count := range cfg.pages {
		fmt.Printf("%s:%d\n", url, count)
	}
}
