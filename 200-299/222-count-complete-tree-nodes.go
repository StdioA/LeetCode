package main

import (
	"fmt"

	"github.com/stdioa/leetcode/bitree"
)

func countNodes(root *bitree.TreeNode) int {
	h := depth(root)
	if h == 0 {
		return 0
	} else {
		if depth(root.Right) == h-1 {
			return (1 << uint(h-1)) + countNodes(root.Right)
		} else {
			return (1 << uint(h-2)) + countNodes(root.Left)
		}
	}
}

func depth(root *bitree.TreeNode) int {
	if root == nil {
		return 0
	}
	dep := 1
	for root.Left != nil {
		dep++
		root = root.Left
	}
	return dep
}

func main() {
	tree, _ := bitree.MakeBiTree("{1,2,3,4,5,6}")
	fmt.Println(countNodes(tree))
}
