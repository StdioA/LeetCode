package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func dfs(node *TreeNode, k int) (amount int, result int) {
	if node == nil {
		// -1 means not found
		return 0, -1
	}

	// Find left subtree
	leftAmount, leftResult := dfs(node.Left, k)
	if leftResult >= 0 {
		return 0, leftResult
	}

	// the node is kth node
	if leftAmount+1 == k {
		return leftAmount + 1, node.Val
	}

	// Find right subtree
	rightAmount, rightResult := dfs(node.Right, k-leftAmount-1)
	if rightResult >= 0 {
		return 0, rightResult
	}

	// Not Found
	return leftAmount + rightAmount + 1, n - 1
}

func kthSmallest(root *TreeNode, k int) int {
	_, result := dfs(root, k)
	return result
}
