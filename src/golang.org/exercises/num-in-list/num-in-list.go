package main

import (
	"fmt"
)

func NumInList(list []int, num int) {
	for _, i := range list {
		if i == num {
			fmt.Printf("True. Value found. \n")
			return
		}
	}
	fmt.Printf("False. Value not found. \n")
	return
}

func main() {
	var list = []int{1, 2, 3, 4}
	NumInList(list, 4)

	var badList = []int{1, 2, 3}
	NumInList(badList, 4)
}
