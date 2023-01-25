package main

import "fmt"

func mySqrt(x int) int {
	lastInt := 0
	for i := 1; i*i <= x; i += 1 {
		if (i * i) <= x {
			lastInt = i
		}
	}

	return lastInt
}

func main() {
	fmt.Println(mySqrt(4))
	fmt.Println(mySqrt(16))
}
