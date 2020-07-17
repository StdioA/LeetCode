package main

import (
	"fmt"

	"github.com/stdioa/leetcode/bitree"
)

type TreeNode = bitree.TreeNode

func addBST(root *TreeNode, base int) int {
	if root == nil {
		return base
	}
	rightBase := addBST(root.Right, base)
	root.Val += rightBase
	leftVal := addBST(root.Left, root.Val)
	return leftVal
}

func convertBST(root *TreeNode) *TreeNode {
	addBST(root, 0)
	return root
}

func main() {
	tree, _ := bitree.New("[5,2,13,1,3,8,15]")
	fmt.Println(convertBST(tree))
}
