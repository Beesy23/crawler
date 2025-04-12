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

	rawURL := args[1]
	fmt.Printf("starting crawl of: %s\n", rawURL)

	html, err := getHTML(rawURL)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Print(html)
}
