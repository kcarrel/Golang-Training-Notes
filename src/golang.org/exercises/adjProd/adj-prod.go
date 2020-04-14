package main

import (
	"fmt"
)

func adjacentElementsProduct(inputArray []int) int {
	var max = -1001
	for i := 0; i < len(inputArray)-1; i++ {
		if inputArray[i]*inputArray[i+1] > max {
			max = inputArray[i] * inputArray[i+1]
		}
	}
	return max
}

func main() {
	var slice1 = []int{3, 6, -2, -5, 7, 3}
	fmt.Println("Expected: 21 ")
	fmt.Printf("Result: %d \n", adjacentElementsProduct(slice1))

	var slice2 = []int{-1, -2}
	fmt.Println("Expected: 2")
	fmt.Printf("Result: %d \n", adjacentElementsProduct(slice2))
}
