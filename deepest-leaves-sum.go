package main

func deepestLeavesSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var (
		result  int
		current = []*TreeNode{root}
	)
	for len(current) > 0 {
		var next []*TreeNode
		result = 0
		for _, node := range current {
			result += node.Val
			if node.Left != nil {
				next = append(next, node.Left)
			}
			if node.Right != nil {
				next = append(next, node.Right)
			}
		}
		current = next
	}
	return result
}
