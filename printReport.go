package main

import (
	"fmt"
	"sort"
)

func printReport(pages map[string]int, baseURL string) {
	fmt.Println("=====")
	fmt.Printf("REPORT for %s\n", baseURL)
	fmt.Println("=====")

	keys := make([]string, 0)
	for k := range pages {
		keys = append(keys, k)
	}

	sort.SliceStable(keys, func(i, j int) bool {
		return pages[keys[i]] >= pages[keys[j]]
	})

	for _, key := range keys {
		fmt.Printf("Found %d internal links to %s\n", pages[key], key)
	}
}
