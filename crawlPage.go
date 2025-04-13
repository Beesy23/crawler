package main

import (
	"fmt"
	"net/url"
)

func crawlPage(rawBaseURL, rawCurrentURL string, pages map[string]int) {

	baseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		fmt.Printf("error parsing baseURL %s: %s\n", rawBaseURL, err)
		return
	}
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

	URLs, err := getURLsFromHTML(HTML, rawBaseURL)
	if err != nil {
		fmt.Printf("error getting urls from currentURL html %s: %s\n", rawCurrentURL, err)
		return
	}
	for _, URL := range URLs {
		crawlPage(rawBaseURL, URL, pages)
	}
}
