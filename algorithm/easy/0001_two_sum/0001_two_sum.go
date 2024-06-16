package easy

func twoSum(nums []int, target int) []int {
	// Save number with index
	numIndexMap := make(map[int]int)

	for i, num := range nums {
		if idx, ok := numIndexMap[target-num]; ok {
			return []int{idx, i}
		}
		numIndexMap[num] = i
	}

	return nil
}
