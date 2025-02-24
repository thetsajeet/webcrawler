package main

import (
	"net/url"
	"strings"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

func GetURLsFromHTML(htmlBody, rawBaseURL string) ([]string, error) {
	htmlReader := strings.NewReader(htmlBody)
	rootNode, err := html.Parse(htmlReader)
	if err != nil {
		return []string{}, err
	}

	urls := []string{}

	baseUrl, err := url.Parse(rawBaseURL)
	if err != nil {
		return []string{}, err
	}

	for n := range rootNode.Descendants() {
		if n.Type == html.ElementNode && n.DataAtom == atom.A {
			for _, attr := range n.Attr {
				if attr.Key == "href" {
					parsedUrl, err := url.Parse(attr.Val)
					if err != nil {
						return []string{}, err
					}

					urls = append(urls, baseUrl.ResolveReference(parsedUrl).String())
				}
			}
		}
	}

	return urls, nil
}
