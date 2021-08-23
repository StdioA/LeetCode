package main

import (
	"github.com/stdioa/leetcode/bitree"
)

type TreeNode = bitree.TreeNode

func buildTree(inorder []int, postorder []int) *TreeNode {
	var (
		length  = len(inorder)
		rootVal = postorder[length-1]
		root    = &TreeNode{Val: rootVal}
		index   = 0
	)
	// Find root index
	for ; index < length; index++ {
		if inorder[index] == rootVal {
			break
		}
	}
	if index > 0 {
		root.Left = buildTree(inorder[:index], postorder[:index])
	}
	if index < length-1 {
		root.Right = buildTree(inorder[index+1:], postorder[index:length])
	}
	return root
}
