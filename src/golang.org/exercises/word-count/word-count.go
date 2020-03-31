package main

import (
	"fmt"
	"strings"
)

func WordCount(s string) map[string]int {
	m := make(map[string]int)
	a := strings.Fields(s)
	for _, v := range a {
		m[v]++
	}
	fmt.Println("map:", m)
	return m
}

func main() {
	WordCount("this this this this this is a string")
}
