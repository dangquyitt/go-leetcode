package main

import (
	"fmt"
)

func canMakeArithmeticProgression(arr []int) bool {
	if len(arr) <= 2 {
		return true
	}
	step := arr[1] - arr[0]
	n := len(arr) - 1
	for i := 1; i < n; i++ {
		if arr[i+1]-arr[i] != step {
			return false
		}
	}
	return true
}
func main() {
	fmt.Println(canMakeArithmeticProgression([]int{3, 5, 1}))
	fmt.Println(canMakeArithmeticProgression([]int{1, 2, 4}))
}
