package main

import (
	"github.com/stdioa/leetcode/bitree"
)

type TreeNode = bitree.TreeNode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	var (
		rootVal = preorder[0]
		node    = &TreeNode{
			Val: rootVal,
		}
		leftLen int
	)
	// Find left & right tree
	for i, val := range inorder {
		if val == rootVal {
			leftLen = i
		}
	}
	node.Left = buildTree(preorder[1:leftLen+1], inorder[:leftLen])
	node.Right = buildTree(preorder[leftLen+1:], inorder[leftLen+1:])
	return node
}
