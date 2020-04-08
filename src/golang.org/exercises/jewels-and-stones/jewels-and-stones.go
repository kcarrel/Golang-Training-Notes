package main

import (
	"fmt"
)

func numJewelsInStones(j string, s string) int {
	jewels := make(map[string]struct{})
	for i := range j {
		jewels[string(j[i])] = struct{}{}
	}

	count := 0
	for i := range s {
		if _, ok := jewels[string(s[i])]; ok {
			count++
		}
	}
	return count
}

func main() {
	var j = "aA"
	var s = "aAAbbbb"
	receivedCount := numJewelsInStones(j, s)
	fmt.Printf("Count: %d \n", receivedCount)
	fmt.Printf("Expected count: 3 \n")

	var j1 = "z"
	var s1 = "ZZ"
	receivedCount2 := numJewelsInStones(j1, s1)
	fmt.Printf("Count: %d \n", receivedCount2)
	fmt.Printf("Expected Count: 0 \n")

}
