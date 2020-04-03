package main

import (
	"fmt"
)

func subtractProductAndSum(n int) int {
	if n <= 0 {
		fmt.Printf("Result: 0 \n")
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
	result := product - sum
	fmt.Printf("Result: %d \n", result)
	return result
}

func main() {
	subtractProductAndSum(234)
	fmt.Printf("Expected result: 15 \n")

	subtractProductAndSum(4421)
	fmt.Printf("Expected result: 21 \n")

	subtractProductAndSum(0)
	fmt.Printf("Expected result: 0 \n")

}
