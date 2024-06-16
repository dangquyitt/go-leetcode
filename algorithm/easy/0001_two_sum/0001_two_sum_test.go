package easy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type (
	input struct {
		nums     []int
		expected int
	}
)

func TestTwoSum(t *testing.T) {
	testCases := []struct {
		name   string
		input  input
		output []int
	}{
		{
			name: "Test case 1",
			input: input{
				nums:     []int{2, 7, 11, 15},
				expected: 9,
			},
			output: []int{0, 1},
		},
		{
			name: "Test case 2",
			input: input{
				nums:     []int{3, 2, 4},
				expected: 6,
			},
			output: []int{1, 2},
		},
		{
			name: "Test case 3",
			input: input{
				nums:     []int{3, 3},
				expected: 6,
			},
			output: []int{0, 1},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := twoSum(tc.input.nums, tc.input.expected)
			assert.EqualValues(t, tc.output, result)
		})
	}
}
