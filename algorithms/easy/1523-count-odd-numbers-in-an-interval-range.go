package main

import (
	"fmt"
	"strings"
)

// 1
//func countOdds(low int, high int) int {
//		result := 0
//	for i := low; i <= high; i++ {
//		if i%2 != 0 {
//			result++
//		}
//	}
//	return result
//}

// 2
func countOdds(low int, high int) int {
	if low%2 == 0 {
		low++
	}
	if low > high {
		return 0
	}
	return (high-low)/2 + 1
}

func main() {
	fmt.Println(strings.ReplaceAll("1523-count-odd-numbers-in-an-interval-range", "-", " "))
	fmt.Println(countOdds(3, 7))
	fmt.Println(countOdds(8, 10))
	fmt.Println(countOdds(11, 11))
	fmt.Println(countOdds(10, 10))
}
