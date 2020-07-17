package main

import (
	"fmt"

	"github.com/stdioa/leetcode/bitree"
)

type TreeNode = bitree.TreeNode

func flattenSub(left, right *TreeNode) *TreeNode {
	var list, p *TreeNode
	if left != nil {
		list = left
		list.Left, list.Right = nil, flattenSub(left.Left, left.Right)
	}
	if right != nil {
		right.Left, right.Right = nil, flattenSub(right.Left, right.Right)
		if list == nil {
			list = right
		} else {
			p = list
			for p.Right != nil {
				p = p.Right
			}
			p.Right = right
		}
	}
	return list
}

func flatten(root *TreeNode) {
	if root != nil {
		root.Right = flattenSub(root.Left, root.Right)
		root.Left = nil
	}
}

func main() {
	tree, _ := bitree.MakeBiTree("{1,2,5,3,4,null,6}")
	flatten(tree)
	fmt.Println(tree)
}
