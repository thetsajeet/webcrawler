package main

import (
	"errors"
	"strconv"
)

type CrawlArgs struct {
	baseURL        string
	maxConcurrency int
	maxPages       int
}

const DEF_MAX_CONCURRENCY int = 1
const DEF_MAX_PAGES int = 20

func GetArgs(args []string) (CrawlArgs, error) {
	if len(args) == 0 {
		return CrawlArgs{}, errors.New(NO_ARG)
	}
	if len(args) >= 4 {
		return CrawlArgs{}, errors.New(TOO_MANY_ARGS)
	}

	crawl := CrawlArgs{
		baseURL:        args[0],
		maxConcurrency: DEF_MAX_CONCURRENCY,
		maxPages:       DEF_MAX_PAGES,
	}

	if 1 < len(args) {
		maxConcurrency, err := strconv.Atoi(args[1])
		if err != nil {
			return crawl, err
		}
		crawl.maxConcurrency = maxConcurrency
	}

	if 2 < len(args) {
		maxPages, err := strconv.Atoi(args[2])
		if err != nil {
			return crawl, err
		}
		crawl.maxPages = maxPages
	}

	return crawl, nil
}
