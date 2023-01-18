package main

import "fmt"

func numWaterBottles(numBottles int, numExchange int) int {
	// First drink
	result := numBottles
	emptyBottles := numBottles

	for emptyBottles >= numExchange {
		// Exchange
		numBottles = emptyBottles / numExchange
		emptyBottles = emptyBottles % numExchange

		// Drink
		emptyBottles += numBottles
		result += numBottles
	}
	return result
}

func main() {
	fmt.Println(numWaterBottles(15, 4))
	fmt.Println(numWaterBottles(9, 3))
}
