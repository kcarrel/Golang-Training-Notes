package main

import (
	"fmt"
)

func RevString(word string) string {
	var rev string
	for _, i := range word {
		rev = string(i) + rev
	}
	fmt.Printf("Forward: %s \n", word)
	fmt.Printf("Reverse: %s  \n", rev)
	return rev
}

func main() {
	RevString("word")
	RevString("another word")
}
