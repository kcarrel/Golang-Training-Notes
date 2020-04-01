package main

import (
	"fmt"
)

func FizzBuzz(n int) {
	for i := 1; i <= n; i++ {
		switch {
		case i%3 == 0 && i%5 == 0:
			fmt.Printf("Fizz Buzz \n")
		case i%3 == 0:
			fmt.Printf("Fizz \n")
		case i%5 == 0:
			fmt.Printf("Buzz \n")
		default:
			fmt.Println(i)
		}
		if i == n {
			fmt.Printf(", ")
		}
	}
	fmt.Println()
}

func main() {
	FizzBuzz(2)
	FizzBuzz(33)
	FizzBuzz(25)

}
