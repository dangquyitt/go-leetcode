package main

import (
	"fmt"
)

func average(salary []int) float64 {
	if len(salary) <= 2 {
		return 0
	}
	total := salary[0]
	indexMaxSalary := 0
	indexMinSalary := 0
	n := len(salary)
	for i := 1; i < n; i++ {
		total += salary[i]
		if salary[indexMaxSalary] < salary[i] {
			indexMaxSalary = i
		}
		if salary[indexMinSalary] > salary[i] {
			indexMinSalary = i
		}
	}
	result := float64(total-salary[indexMinSalary]-salary[indexMaxSalary]) / float64(n-2)
	return result
}

func main() {
	fmt.Println(average([]int{4000, 3000, 1000, 2000}))
	fmt.Println(average([]int{1000, 2000, 3000}))
}
