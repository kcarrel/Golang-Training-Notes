package main

import (
	"fmt"
)

func GCD(a, b int) {
	for b != 0 {
		a, b = b, a%b
	}
	fmt.Printf("GCD: %d \n", a)
}

func main() {
	GCD(10, 5)
	GCD(25, 5)
	GCD(100, 9)
}
