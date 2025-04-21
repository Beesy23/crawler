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

	for url, linkCount := range pages {
		sortedPages = append(sortedPages, page{url, linkCount})
	}

	slices.SortFunc(sortedPages, func(a, b page) int {
		if a.numberLinks != b.numberLinks {
			return b.numberLinks - a.numberLinks
		}

		if a.numberLinks < b.numberLinks {
			return -1
		} else if a.pageName > b.pageName {
			return 1
		}
		return 0
	})

	return sortedPages
}
