package main

import "strconv"

func tree2str(root *TreeNode) string {
	if root == nil {
		return ""
	}
	var s = strconv.Itoa(root.Val)
	left, right := tree2str(root.Left), tree2str(root.Right)
	if left != "" {
		s += "(" + left + ")"
	}
	if right != "" {
		if left == "" {
			s += "()"
		}
		s += "(" + right + ")"
	}
	return s
}
