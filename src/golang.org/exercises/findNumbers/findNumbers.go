package main

import (
	"fmt"
)

func findNumbers(nums []int) int {
	count := 0
	for _, num := range nums {
		digits := 1
		for num > 9 {
			digits++
			num /= 10
		}
		if digits%2 == 0 {
			count++
		}
	}
	fmt.Printf("Count: %d \n", count)
	return count
}

func main() {
	var expectedList = []int{10, 2, 300, 4}
	findNumbers(expectedList)
	fmt.Printf("Expected count: 1 \n")

	var emptyList = []int{}
	findNumbers(emptyList)
	fmt.Printf("Expected count: 0 \n")

	var zeroList = []int{1, 2, 3, 4}
	findNumbers(zeroList)
	fmt.Printf("Expected count: 0 \n")

}
