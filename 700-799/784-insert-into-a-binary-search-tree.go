package main

import (
	"fmt"

	"github.com/stdioa/leetcode/bitree"
)

func insertIntoBST(root *bitree.TreeNode, val int) *bitree.TreeNode {
	if root == nil {
		return &bitree.TreeNode{Val: val}
	}
	node := root
	for {
		if val < node.Val {
			if node.Left == nil {
				node.Left = &bitree.TreeNode{Val: val}
				break
			} else {
				node = node.Left
			}
		} else {
			if node.Right == nil {
				node.Right = &bitree.TreeNode{Val: val}
				break
			} else {
				node = node.Right
			}
		}
	}
	return root
}

func main() {
	tree, _ := bitree.MakeBiTree("{4,2,7,1,3}")
	tree = insertIntoBST(tree, 5)
	fmt.Println(tree)
}
