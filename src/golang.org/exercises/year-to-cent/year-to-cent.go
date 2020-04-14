package main

import (
	"fmt"
)

func centuryFromYear(year int) int {
	return ((year - 1) / 100) + 1
}

func main() {
	fmt.Printf("Expected: 20 \n")
	fmt.Printf("Result: %d \n", centuryFromYear(1905))

	fmt.Printf("Expected: 17 \n")
	fmt.Printf("Result: %d \n", centuryFromYear(1700))

	fmt.Printf("Expected: 15 \n")
	fmt.Printf("Result: %d \n", centuryFromYear(1500))
}
