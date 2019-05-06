package main

import (
	"fmt"

	"github.com/stdioa/leetcode/bitree"
)

type TreeNode = bitree.TreeNode

func flipEquiv(root1 *TreeNode, root2 *TreeNode) bool {
	if (root1 == nil) != (root2 == nil) {
		return false
	} else if root1 == nil && root2 == nil {
		return true
	} else if root1.Val != root2.Val {
		return false
	}

	l1, r1 := root1.Left, root1.Right
	l2, r2 := root2.Left, root2.Right

	if r1 == nil {
		l1, r1 = r1, l1
	} else if l1 != nil && l1.Val > r1.Val {
		l1, r1 = r1, l1
	}
	if r2 == nil {
		l2, r2 = r2, l2
	} else if l2 != nil && l2.Val > r2.Val {
		l2, r2 = r2, l2
	}
	return flipEquiv(l1, l2) && flipEquiv(r1, r2)
}

func main() {
	tree1, _ := bitree.MakeBiTree("{1,2,3,4,5,6,null,null,null,7,8}")
	tree2, _ := bitree.MakeBiTree("{1,3,2,null,6,4,5,null,null,null,null,8,7}")
	fmt.Println(flipEquiv(tree1, tree2))
}
