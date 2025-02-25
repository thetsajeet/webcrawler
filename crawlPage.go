package main

import (
	"fmt"
	"log"
	"net/url"
)

func (cfg *config) CrawlPage(rawCurrentURL string) {
	cfg.concurrencyControl <- struct{}{}
	defer func() {
		<-cfg.concurrencyControl
		defer cfg.wg.Done()
	}()

	if cfg.hasReachedLimit() {
		return
	}

	parsedCurrentURL, err := url.Parse(rawCurrentURL)
	if err != nil {
		fmt.Println(err)
		return
	}

	if cfg.baseURL.Host != parsedCurrentURL.Host {
		return
	}

	nURL, err := NormalizeURL(rawCurrentURL)
	if err != nil {
		fmt.Println(err)
		return
	}

	if isFirst := cfg.addPageVisit(nURL); !isFirst {
		return
	}

	log.Default().Printf("Crawling %s\n", rawCurrentURL)

	htmlContent, err := GetHtml(rawCurrentURL)
	if err != nil {
		fmt.Println(err)
		return
	}

	urls, err := GetURLsFromHTML(htmlContent, cfg.baseURL.String())
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, rawURL := range urls {
		cfg.wg.Add(1)
		go cfg.CrawlPage(rawURL)
	}

}
