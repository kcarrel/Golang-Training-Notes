package main

import (
	"fmt"
)

func checkPalindrome(inputString string) bool {
	for i := 0; i < len(inputString)/2; i++ {
		if inputString[i] != inputString[len(inputString)-i-1] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println("Expected: true")
	fmt.Printf("Result: %t \n", checkPalindrome("anna"))

	fmt.Println("Expected: false")
	fmt.Printf("Result: %t \n", checkPalindrome("banana"))
}
