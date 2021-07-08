package main

import "github.com/stdioa/leetcode/bitree"

func lowestCommonAncestor(root, p, q *bitree.TreeNode) *bitree.TreeNode {
	if root == nil {
		return root
	}
	if p.Val > q.Val {
		p, q = q, p
	}
	if p.Val <= root.Val && root.Val <= q.Val {
		return root
	}
	if root.Val > q.Val {
		return lowestCommonAncestor(root.Left, p, q)
	}
	return lowestCommonAncestor(root.Right, p, q)
}
