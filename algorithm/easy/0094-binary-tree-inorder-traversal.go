package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	var nums = []int{}
	if root == nil {
		return nums
	}

	nums = append(nums, inorderTraversal(root.Left)...)
	nums = append(nums, root.Val)
	nums = append(nums, inorderTraversal(root.Right)...)
	return nums
}

func main() {
	var root TreeNode
	root.Val = 1
	root.Left = &TreeNode{
		Val: 3,
	}
	root.Right = &TreeNode{
		Val: 4,
	}
	fmt.Printf("%v", inorderTraversal(&root))
}
