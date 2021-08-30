package main

import (
	"github.com/stdioa/leetcode/bitree"
)

type TreeNode = bitree.TreeNode

func dfs(root *TreeNode, k int) (result *int, rank int) {
	if root == nil {
		return
	}
	rightResult, rightRank := dfs(root.Right, k)
	if rightResult != nil {
		result = rightResult
		return
	}
	if k-rightRank == 1 {
		result = &root.Val
		return
	}
	leftResult, leftRank := dfs(root.Left, k-rightRank-1)
	return leftResult, leftRank + rightRank + 1
}

func kthLargest(root *TreeNode, k int) int {
	result, _ := dfs(root, k)
	if result != nil {
		return *result
	}
	return 0
}
