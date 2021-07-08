package main

import (
	"fmt"

	"github.com/stdioa/leetcode/bitree"
)

func rangeSumBST(root *bitree.TreeNode, L int, R int) int {
	if root == nil {
		return 0
	}
	result := 0
	if L <= root.Val && root.Val <= R {
		result += root.Val
	}
	if root.Val > L {
		result += rangeSumBST(root.Left, L, R)
	}
	if root.Val < R {
		result += rangeSumBST(root.Right, L, R)
	}
	return result
}

func main() {
	tree, _ := bitree.MakeBiTree("{10,5,15,3,7,13,18,1,null,6}")
	fmt.Println(rangeSumBST(tree, 6, 10))
}
