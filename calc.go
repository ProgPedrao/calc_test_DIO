package main

import (
	"fmt"
)

func main() {

	sum := sum(3, 2, 1)
	subtraction := subtraction(3, 2)
	multiplication := multiplication(3, 2, 1)
	division := division(3, 2)

	fmt.Printf("\nsum - %d", sum)
	fmt.Printf("\nsubtraction - %d", subtraction)
	fmt.Printf("\nmultiplication - %d", multiplication)
	fmt.Printf("\ndivision - %f", division)
}

func sum(i ...int) int {
	result := 0

	for _, value := range i {
		result = result + value
	}

	return result
}

func subtraction(totalValue int, subtractValue int) int {
	result := totalValue - subtractValue

	return result
}

func multiplication(i ...int) int {
	result := 1

	for _, value := range i {
		result = result * value
	}

	return result
}

func division(totalValue int, subtractValue int) float64 {
	result := float64(totalValue) / float64(subtractValue)

	return result
}
