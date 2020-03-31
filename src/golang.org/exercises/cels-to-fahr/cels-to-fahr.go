package main

import "fmt"

func squares() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}

func main() {
	var fahr int
	var celsius int
	lower := 0
	upper := 300
	step := 20

	celsius = lower
	for celsius <= upper {
		fahr = (9/5)*celsius + 32
		fmt.Printf("Celsius: %d Fahr: %d\n", celsius, fahr)
		celsius = celsius + step
	}

}
