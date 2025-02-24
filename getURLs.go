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

	for n := range rootNode.Descendants() {
		if n.Type == html.ElementNode && n.DataAtom == atom.A {
			for _, attr := range n.Attr {
				if attr.Key == "href" {
					if parsedUrl, err := url.Parse(attr.Val); err != nil {
						return []string{}, err
					} else if parsedUrl.IsAbs() {
						urls = append(urls, parsedUrl.String())
					} else {
						urls = append(urls, rawBaseURL+parsedUrl.String())
					}
				}
			}
		}
	}

	return urls, nil
}
