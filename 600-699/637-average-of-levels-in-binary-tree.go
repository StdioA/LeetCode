package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func averageOfLevels(root *TreeNode) []float64 {
	var (
		layer   = []*TreeNode{root}
		answers []float64
	)
	for len(layer) > 0 {
		var (
			nextLayer    []*TreeNode
			total, count int
		)
		for _, node := range layer {
			total += node.Val
			count++
			if node.Left != nil {
				nextLayer = append(nextLayer, node.Left)
			}
			if node.Right != nil {
				nextLayer = append(nextLayer, node.Right)
			}
		}
		layer = nextLayer
		answers = append(answers, float64(total)/float64(count))
	}
	return answers
}
