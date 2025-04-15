package main

import (
	"fmt"
	"net/url"
	"sync"
)

type config struct {
	pages              map[string]int
	baseURL            *url.URL
	mu                 *sync.Mutex
	concurrencyControl chan struct{}
	wg                 *sync.WaitGroup
}

func (cfg *config) crawlPage(rawCurrentURL string) {
	baseURL := cfg.baseURL
	pages := cfg.pages

	currentURL, err := url.Parse(rawCurrentURL)
	if err != nil {
		fmt.Printf("error parsing currentURL %s: %s\n", rawCurrentURL, err)
		return
	}
	if baseURL.Host != currentURL.Host {
		return
	}

	normalisedURL, err := normalizeURL(rawCurrentURL)
	if err != nil {
		fmt.Printf("error normalising currentURL %s: %s\n", rawCurrentURL, err)
		return
	}

	if _, visited := pages[normalisedURL]; visited {
		pages[normalisedURL]++
		return
	} else {
		pages[normalisedURL] = 1
	}

	HTML, err := getHTML(rawCurrentURL)
	if err != nil {
		fmt.Printf("error retrieving html from currentURL %s: %s\n", rawCurrentURL, err)
		return
	}
	fmt.Printf("Crawler succesfully checked %s\n", rawCurrentURL)

	URLs, err := getURLsFromHTML(HTML, baseURL.String())
	if err != nil {
		fmt.Printf("error getting urls from currentURL html %s: %s\n", rawCurrentURL, err)
		return
	}
	for _, URL := range URLs {
		cfg.crawlPage(URL)
	}
}
