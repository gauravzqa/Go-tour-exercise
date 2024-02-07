package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	m := make(map[string]int)

	for _, val := range strings.Fields(s) {
		m[val]++
	}

	return m
}

func main() {
	wc.Test(WordCount)
}
