package main

import (
	"fmt"

	"github.com/stdioa/leetcode/bitree"
)

type TreeNode = bitree.TreeNode

func calc(tree *TreeNode) (int, int) {
	if tree == nil {
		return 0, 0
	}
	lDepth, lDiam := calc(tree.Left)
	rDepth, rDiam := calc(tree.Right)
	maxDiam := lDepth + rDepth
	if lDiam > maxDiam {
		maxDiam = lDiam
	}
	if rDiam > maxDiam {
		maxDiam = rDiam
	}

	if lDepth < rDepth {
		lDepth = rDepth
	}
	return lDepth + 1, maxDiam
}

func diameterOfBinaryTree(root *TreeNode) int {
	_, diameter := calc(root)
	return diameter
}

func main() {
	// tree, _ := bitree.MakeBiTree("[1,2,3,4,5]")
	tree, _ := bitree.MakeBiTree("[1,2,#,4,5,6,#,7,8]")
	fmt.Println(diameterOfBinaryTree(tree))
}
