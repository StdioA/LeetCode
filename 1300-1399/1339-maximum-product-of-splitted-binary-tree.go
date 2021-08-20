package main

import (
	"fmt"

	"github.com/stdioa/leetcode/bitree"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode = bitree.TreeNode

func calSum(node *TreeNode) int {
	if node == nil {
		return 0
	}
	node.Val += calSum(node.Left) + calSum(node.Right)
	return node.Val
}

func absDis(a, b int) int {
	dis := a - b
	if dis < 0 {
		dis = -dis
	}
	return dis
}

func findMin(node *TreeNode, target int) int {
	val := node.Val
	dis := absDis(val, target)
	if dis == 0 {
		return val
	}
	if node.Left != nil {
		leftVal := findMin(node.Left, target)
		leftDis := absDis(leftVal, target)
		if leftDis == 0 {
			return leftVal
		} else if leftDis < dis {
			dis = leftDis
			val = leftVal
		}
	}
	if node.Right != nil {
		rightVal := findMin(node.Right, target)
		rightDis := absDis(rightVal, target)
		if rightDis == 0 {
			return rightVal
		} else if rightDis < dis {
			dis = rightDis
			val = rightVal
		}
	}
	return val
}

func maxProduct(root *TreeNode) int {
	const mod = 1000000007
	var (
		total  = calSum(root)
		target = total / 2
	)
	val := findMin(root, target)
	return (val % mod) * ((total - val) % mod) % mod
}

func main() {
	tree, _ := bitree.MakeBiTree("[1,null,2,3,4,null,null,5,6]")
	fmt.Println(maxProduct(tree))
}
