package main

import (
	"fmt"
	// "code.google.com/p/go-tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	fields := strings.Fields(s)
	var m map[string]int
	m = make(map[string]int)

	for _, f := range fields {
		m[f] += 1
	}
	return m
}

func main() {
	m := WordCount("aaa bbb aaa dd")
	fmt.Println(m)
}
