package main

import (
	"fmt"
	"net/url"
	"os"
	"strconv"
	"sync"
)

func main() {
	args := os.Args

	if len(args) < 4 {
		fmt.Print("not enough arguments provided\n")
		os.Exit(1)
	}
	if len(args) > 4 {
		fmt.Print("too many arguments provided\n")
		os.Exit(1)
	}

	rawBaseURL := args[1]
	maxConcurrecy, err := strconv.Atoi(args[2])
	if err != nil {
		fmt.Printf("Error max concurrency not an integer: %v\n", err)
		return
	}
	maxPages, err := strconv.Atoi(args[3])
	if err != nil {
		fmt.Printf("Error max pages not an integer: %v\n", err)
		return
	}
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
		concurrencyControl: make(chan struct{}, maxConcurrecy),
		wg:                 &sync.WaitGroup{},
		maxPages:           maxPages,
	}

	cfg.crawlPage(rawBaseURL)
	cfg.wg.Wait()

	fmt.Println("\nCrawling results:")
	for url, count := range cfg.pages {
		fmt.Printf("%s:%d\n", url, count)
	}
	printReport(cfg.pages, rawBaseURL)
}
