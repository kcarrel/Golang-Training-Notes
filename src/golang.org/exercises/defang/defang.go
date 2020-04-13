package main

import (
	"fmt"
)

func defangIPaddr(address string) string {
	result := ""
	for _, char := range address {
		if string(char) == "." {
			result += "[.]"
		} else {
			result += string(char)
		}
	}
	return result
}

func main() {
	input1 := "1.1.1.1"
	result1 := defangIPaddr(input1)
	fmt.Printf("Expected: 1[.]1[.]1[.]1 \n")
	fmt.Printf("Result:  %s \n", result1)

	input2 := "255.100.50.0"
	result2 := defangIPaddr(input2)
	fmt.Printf("Expected: 255[.]100[.]50[.]0 \n")
	fmt.Printf("Result:  %s \n", result2)

}
