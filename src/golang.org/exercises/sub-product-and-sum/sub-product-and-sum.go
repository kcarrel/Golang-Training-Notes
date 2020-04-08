package main

import (
	"fmt"
)

func subtractProductAndSum(n int) int {
	if n <= 0 {
		return 0
	}

	sum := 0
	product := 1

	for n > 0 {
		mod := n % 10
		sum += mod
		product *= mod
		n /= 10
	}
	return product - sum
}

func main() {
	result := subtractProductAndSum(234)
	fmt.Printf("Result: %d \n", result)
	fmt.Printf("Expected result: 15 \n")

	result1 := subtractProductAndSum(4421)
	fmt.Printf("Result: %d \n", result1)
	fmt.Printf("Expected result: 21 \n")

	result2 := subtractProductAndSum(0)
	fmt.Printf("Result: %d \n", result2)
	fmt.Printf("Expected result: 0 \n")
}
