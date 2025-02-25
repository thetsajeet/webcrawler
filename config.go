package main

import (
	"net/url"
	"sync"
)

type config struct {
	pages              map[string]int
	baseURL            *url.URL
	mu                 *sync.Mutex
	concurrencyControl chan struct{}
	wg                 *sync.WaitGroup
	maxPages           int
}

func (cfg *config) addPageVisit(nURL string) (isFirst bool) {
	cfg.mu.Lock()
	defer cfg.mu.Unlock()

	if _, exists := cfg.pages[nURL]; exists {
		cfg.pages[nURL]++
		return false
	}

	cfg.pages[nURL] = 1
	return true
}

func createConfig(rawBaseUrl string, maxConcurrency, maxPages int) (*config, error) {
	baseURL, err := url.Parse(rawBaseUrl)
	if err != nil {
		return nil, err
	}

	return &config{
		pages:              map[string]int{},
		baseURL:            baseURL,
		concurrencyControl: make(chan struct{}, maxConcurrency),
		wg:                 &sync.WaitGroup{},
		mu:                 &sync.Mutex{},
		maxPages:           maxPages,
	}, nil
}

func (cfg *config) hasReachedLimit() bool {
	cfg.mu.Lock()
	defer cfg.mu.Unlock()

	return len(cfg.pages) >= cfg.maxPages
}
