package main

import (
	"net/url"
	"strings"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

func getURLsFromHTML(htmlBody, rawBaseURL string) ([]string, error) {
	urls := []string{}
	baseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		return []string{}, err
	}

	node, err := html.Parse(strings.NewReader(htmlBody))
	if err != nil {
		return []string{}, err
	}

	for n := range node.Descendants() {
		if n.Type == html.ElementNode && n.DataAtom == atom.A {
			for _, a := range n.Attr {
				if a.Key == "href" {
					parsedURL, err := url.Parse(a.Val)
					if err != nil {
						continue
					}

					resolvedURL := baseURL.ResolveReference(parsedURL)
					urls = append(urls, resolvedURL.String())
				}
			}
		}
	}
	return urls, nil
}
