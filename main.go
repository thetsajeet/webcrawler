package main

import (
	"fmt"
	"os"
	"time"
)

const NO_ARG = "no website provided"
const TOO_MANY_ARGS = "too many arguments provided"

func main() {
	base_url, err := GetArgs(os.Args[1:])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("starting crawl of: %s\n", base_url)

	cfg, err := createConfig(base_url, 5)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	now := time.Now()
	cfg.wg.Add(1)
	go cfg.CrawlPage(base_url)
	cfg.wg.Wait()
	end := time.Now()

	for k, v := range cfg.pages {
		fmt.Printf("%s: %d\n", k, v)
	}

	fmt.Printf("Took %v in total\n", end.Sub(now))
}
