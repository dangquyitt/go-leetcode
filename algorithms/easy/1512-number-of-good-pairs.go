package main

import "fmt"

func numIdenticalPairs(nums []int) int {
	if len(nums) < 2 {
		return 0
	}

	lenghNums := len(nums)
	result := 0

	for i := 0; i < lenghNums-1; i++ {
		for j := i + 1; j < lenghNums; j++ {
			if nums[i] == nums[j] {
				result += 1
			}
		}
	}

	return result
}

func main() {
	fmt.Println(numIdenticalPairs([]int{1, 2, 3, 1, 1, 3}))
	fmt.Println(numIdenticalPairs([]int{1, 1, 1, 1}))
}
