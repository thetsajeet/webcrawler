package main

import (
	"net/url"
	"strings"
)

func NormalizeURL(requestURL string) (string, error) {
	parsedRequestURL, err := url.ParseRequestURI(requestURL)
	if err != nil {
		return "", err
	}
	normalisedURL := parsedRequestURL.Host + parsedRequestURL.Path
	normalisedURL = strings.TrimSuffix(normalisedURL, "/")
	return normalisedURL, nil
}
