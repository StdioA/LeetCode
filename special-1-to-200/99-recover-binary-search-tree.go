package main

import (
	"github.com/stdioa/leetcode/bitree"
)

type TreeNode = bitree.TreeNode

func recoverTree(root *TreeNode) {
	// 对 BST 进行中序遍历，并找到两个连续的且不符合升序要求的数对
	// 第一对中的第一个值，和第二对中的第二个值需要互换
	// 注意，这两对可能是同一对
	// 举例：[3,1,4,null,null,2] -> 前序 [1 3 2 4] -> [3 2]
	// 举例：前序 [6, 3, 4, 5, 2] -> [6 3] 和 [5 2]

	var firstElement, secondElement, prevElement *TreeNode
	prevElement = &TreeNode{Val: -2 << 32} // 当前遍历节点的上一个节点

	var traverse func(root *TreeNode)
	traverse = func(root *TreeNode) {
		if root == nil {
			return
		}
		traverse(root.Left)
		// 找到第一个逆序对，如果前一个值更大，那它就应该是第一个有问题的值
		if firstElement == nil && prevElement.Val >= root.Val {
			firstElement = prevElement
		}
		// 找到第二个逆序对，此时有问题的值是第二个值，也就是 root
		if firstElement != nil && prevElement.Val >= root.Val {
			secondElement = root
		}
		prevElement = root
		traverse(root.Right)
	}
	traverse(root)

	firstElement.Val, secondElement.Val = secondElement.Val, firstElement.Val
}
