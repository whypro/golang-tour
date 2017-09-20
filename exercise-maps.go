package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	wordCount := make(map[string]int)
	wordSlice := strings.Fields(s)
	for _, v := range wordSlice {
		wordCount[v] += 1
	}
	return wordCount
}

func main() {
	wc.Test(WordCount)
}
