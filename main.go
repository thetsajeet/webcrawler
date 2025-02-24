package main

import (
	"fmt"
	"os"
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
	pages := map[string]int{}
	CrawlPage(base_url, base_url, pages)
	for k, v := range pages {
		fmt.Printf("%s: %d\n", k, v)
	}
}
