package main

import (
	"fmt"

	"github.com/stdioa/leetcode/bitree"
)

type TreeNode = bitree.TreeNode

func addOneRow(root *TreeNode, val int, depth int) *TreeNode {
	if depth == 1 {
		node := &TreeNode{Val: val}
		node.Left = root
		return node
	}

	// Get all nodes at level (depth - 1)
	var (
		nodes = []*TreeNode{root}
	)
	for i := 1; i < depth-1; i++ {
		var children []*TreeNode
		for _, node := range nodes {
			if node.Left != nil {
				children = append(children, node.Left)
			}
			if node.Right != nil {
				children = append(children, node.Right)
			}
			nodes = children
		}
	}

	// Insert nodes
	for _, node := range nodes {
		left, right := &TreeNode{Val: val}, &TreeNode{Val: val}
		node.Left, left.Left = left, node.Left
		node.Right, right.Right = right, node.Right
	}
	return root
}

func main() {
	tree, _ := bitree.MakeBiTree("[4,2,6,3,1,5]")
	fmt.Println(addOneRow(tree, 1, 2))
}
