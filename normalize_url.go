package main

import (
	"net/url"
	"strings"
)

func normalizeURL(requestURL string) (string, error) {
	parsedRequestURL, err := url.ParseRequestURI(requestURL)
	if err != nil {
		return "", err
	}
	normalisedURL := parsedRequestURL.Host + parsedRequestURL.Port() + parsedRequestURL.RequestURI()
	normalisedURL = strings.TrimSuffix(normalisedURL, "/")
	return normalisedURL, nil
}
