package main

import (
	"fmt"

	"github.com/stdioa/leetcode/bitree"
)

func inorderTraversal(root *bitree.TreeNode) {
	if root == nil {
		return
	}

	var (
		stack = []*bitree.TreeNode{}
		p     = root
	)
	for p != nil || len(stack) > 0 {
		for p != nil {
			stack = append(stack, p)
			p = p.Left
		}
		if len(stack) > 0 {
			p = stack[len(stack)-1]
			fmt.Println(p.Val)
			stack = stack[:len(stack)-1]
			p = p.Right
		}
	}
}

func main() {
	tree, _ := bitree.MakeBiTree("{1,2,3,4,5,6,7}")
	inorderTraversal(tree)
}
