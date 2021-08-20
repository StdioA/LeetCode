package main

import (
	"github.com/stdioa/leetcode/bitree"
)

type TreeNode = bitree.TreeNode

func dfs(node *TreeNode, max int) int {
	if node == nil {
		return 0
	}
	var result int
	if node.Val >= max {
		result = 1
	}
	if node.Val > max {
		max = node.Val
	}
	return result + dfs(node.Left, max) + dfs(node.Right, max)
}

func goodNodes(root *TreeNode) int {
	return dfs(root, -10001)
}
