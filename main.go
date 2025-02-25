package main

import (
	"fmt"
	"os"
	"time"
)

const NO_ARG = "no website provided"
const TOO_MANY_ARGS = "too many arguments provided"

func main() {
	crawlArgs, err := GetArgs(os.Args[1:])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	cfg, err := createConfig(crawlArgs.baseURL, crawlArgs.maxConcurrency, crawlArgs.maxPages)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	now := time.Now()
	cfg.wg.Add(1)
	go cfg.CrawlPage(crawlArgs.baseURL)
	cfg.wg.Wait()
	end := time.Now()

	printReport(cfg.pages, crawlArgs.baseURL)

	fmt.Printf("Took %v in total\n", end.Sub(now))
}
