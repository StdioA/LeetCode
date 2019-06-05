package main

import (
	"fmt"

	"github.com/stdioa/leetcode/bitree"
)

type TreeNode = bitree.TreeNode

func bstFromPreorder(preorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	val := preorder[0]
	tree := &TreeNode{Val: val}
	i := 1
	for i < len(preorder) {
		if preorder[i] > val {
			break
		}
		i++
	}
	tree.Left = bstFromPreorder(preorder[1:i])
	tree.Right = bstFromPreorder(preorder[i:])
	return tree
}

func main() {
	fmt.Println(bstFromPreorder([]int{8, 5, 1, 7, 10, 12}))
}
