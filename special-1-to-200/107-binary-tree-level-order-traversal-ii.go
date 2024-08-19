package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	// 层序遍历完，倒过来就行了
	var (
		results [][]int
		level   = []*TreeNode{root}
	)
	for len(level) > 0 {
		var (
			levelResult = make([]int, 0, len(level))
			nextLevel   []*TreeNode
		)
		for _, node := range level {
			levelResult = append(levelResult, node.Val)
			if node.Left != nil {
				nextLevel = append(nextLevel, node.Left)
			}
			if node.Right != nil {
				nextLevel = append(nextLevel, node.Right)
			}
		}
		level = nextLevel
		results = append(results, levelResult)
	}
	var levels = len(results)
	for i := 0; i < levels/2; i++ {
		results[i], results[levels-1-i] = results[levels-1-i], results[i]
	}
	return results
}
