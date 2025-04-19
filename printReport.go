package main

import (
	"fmt"
	"slices"
)

type page struct {
	pageName    string
	numberLinks int
}

func printReport(pages map[string]int, baseURL string) {
	fmt.Println("=============================")
	fmt.Printf("REPORT for %s\n", baseURL)
	fmt.Println("=============================")

	sortedPages := sortPages(pages)
	for _, page := range sortedPages {
		fmt.Printf("Found %d internal links to %s\n", page.numberLinks, page.pageName)
	}
}

func sortPages(pages map[string]int) []page {
	sortedPages := []page{}
	for url, timeVisited := range pages {
		pageToSort := page{url, timeVisited}
		if slices.Contains(sortedPages, pageToSort) {

		}
	}
	return sortedPages
}
