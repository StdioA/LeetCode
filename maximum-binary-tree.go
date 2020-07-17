package main

import (
	"fmt"

	"github.com/stdioa/leetcode/bitree"
)

type TreeNode = bitree.TreeNode

func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	maxi, max := nums[0], 0
	for i, n := range nums {
		if n > max {
			maxi, max = i, n
		}
	}
	tree := &TreeNode{Val: max}
	tree.Left = constructMaximumBinaryTree(nums[:maxi])
	tree.Right = constructMaximumBinaryTree(nums[maxi+1:])
	return tree
}

func main() {
	fmt.Println(constructMaximumBinaryTree([]int{3, 2, 1, 6, 0, 5}))
}
