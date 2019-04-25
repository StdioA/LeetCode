package main

import (
	"fmt"

	"github.com/stdioa/leetcode/bitree"
)

func isValidNode(root *TreeNode, min, max int) bool {
	if root == nil {
		return true
	}
	if !(min < root.Val && root.Val < max) {
		return false
	}
	return isValidNode(root.Left, min, root.Val) && isValidNode(root.Right, root.Val, max)
}

func isValidBST(root *TreeNode) bool {
	return isValidNode(root, -2147483648000, 2147483647000)
}

func main() {
	tree, _ := bitree.MakeBiTree("{10,5,15,null,null,6,20}")
	fmt.Println(isValidBST(tree))
}
