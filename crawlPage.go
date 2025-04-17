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

	isFirst := cfg.addPageVisit(normalisedURL)
	if !isFirst {
		return
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
		cfg.wg.Add(1)
		go func(urlToCrawl string) {
			cfg.concurrencyControl <- struct{}{}
			defer func() {
				<-cfg.concurrencyControl
				cfg.wg.Done()
			}()

			cfg.crawlPage(urlToCrawl)
		}(URL)
	}
}

func (cfg *config) addPageVisit(normalizedURL string) (isFirst bool) {
	cfg.mu.Lock()
	defer cfg.mu.Unlock()

	if _, ok := cfg.pages[normalizedURL]; ok {

		return false
	}

	cfg.pages[normalizedURL] = 1
	return true
}
