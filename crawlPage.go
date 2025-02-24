package main

import (
	"fmt"
	"log"
	"net/url"
)

func CrawlPage(rawBaseURL, rawCurrentURL string, pages map[string]int) {
	parsedBaseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		fmt.Println(err)
		return
	}

	parsedCurrentURL, err := url.Parse(rawCurrentURL)
	if err != nil {
		fmt.Println(err)
		return
	}

	if parsedBaseURL.Host != parsedCurrentURL.Host {
		return
	}

	nURL, err := NormalizeURL(rawCurrentURL)
	if err != nil {
		fmt.Println(err)
		return
	}

	if _, exists := pages[nURL]; exists {
		pages[nURL] += 1
		return
	}

	pages[nURL] = 1

	log.Default().Printf("Crawling %s\n", rawCurrentURL)

	htmlContent, err := GetHtml(rawCurrentURL)
	if err != nil {
		fmt.Println(err)
		return
	}

	urls, err := GetURLsFromHTML(htmlContent, rawBaseURL)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, rawURL := range urls {
		CrawlPage(rawBaseURL, rawURL, pages)
	}
}
