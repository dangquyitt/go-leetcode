package main

import (
	"fmt"
	"math/rand"
)

type Solution struct {
	w []int
}

func Constructor(w []int) Solution {
	return Solution{
		w: w,
	}
}

func (this *Solution) PickIndex() int {
	sumOfWeight := 0
	for _, v := range this.w {
		sumOfWeight += v
	}

	randVal := rand.Intn(sumOfWeight)
	for i, v := range this.w {
		if randVal < v {
			return i
		}
		randVal -= v
	}

	return 0
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(w);
 * param_1 := obj.PickIndex();
 */
func main() {
	obj := Constructor([]int{1, 4, 3, 5})
	for i := 0; i < 100; i++ {
		fmt.Println(obj.PickIndex())
	}
}
