package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

const BAD_RESPONSE = "got a bad response"
const BAD_CONTENT_TYPE = "invalid content type"

func GetHtml(rawURL string) (string, error) {
	resp, err := http.Get(rawURL)
	if err != nil {
		return "", err
	}

	if resp.StatusCode >= 400 {
		return "", fmt.Errorf("%s: %d", BAD_RESPONSE, resp.StatusCode)
	}

	if contentType := resp.Header.Get("content-type"); !strings.Contains(contentType, "text/html") {
		return "", fmt.Errorf("%s: %s", BAD_CONTENT_TYPE, contentType)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
