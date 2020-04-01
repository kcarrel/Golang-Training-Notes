package main

import (
	"fmt"
)

func SumNum(list []int) {
	var sum int
	for _, i := range list {
		sum += i
	}
	fmt.Printf("Sum: %d \n", sum)
	return
}

func main() {
	var list = []int{1, 2, 3, 4}
	SumNum(list)

	var list2 = []int{1, 2, 3, 4, 10, 21}
	SumNum(list2)
}
