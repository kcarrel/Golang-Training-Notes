package main

import (
	"fmt"
)

func numJewelsInStones(J string, S string) int {
	jewels := make(map[string]struct{})
	for i := range J {
		jewels[string(J[i])] = struct{}{}
	}

	count := 0
	for i := range S {
		if _, ok := jewels[string(S[i])]; ok {
			count++
		}
	}
	fmt.Printf("Count: %d \n", count)
	return count
}

func main() {
	var J = "aA"
	var S = "aAAbbbb"
	numJewelsInStones(J, S)
	fmt.Printf("Expected count: 3 \n")

	var J1 = "z"
	var S1 = "ZZ"
	numJewelsInStones(J1, S1)
	fmt.Printf("Expected Count: 0 \n")

}
