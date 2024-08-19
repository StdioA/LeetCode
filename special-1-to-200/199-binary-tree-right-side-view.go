package main

// import "github.com/stdioa/leetcode/bitree"

// type TreeNode = bitree.TreeNode

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	// 层序遍历的加工
	var (
		level   = []*TreeNode{root}
		results []int
	)
	for len(level) > 0 {
		var (
			nextLevel []*TreeNode
			result    int
		)
		for _, node := range level {
			result = node.Val
			if node.Left != nil {
				nextLevel = append(nextLevel, node.Left)
			}
			if node.Right != nil {
				nextLevel = append(nextLevel, node.Right)
			}
		}
		results = append(results, result)
		level = nextLevel
	}
	return results
}
