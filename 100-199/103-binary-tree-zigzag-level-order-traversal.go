/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var (
		nodes  = []*TreeNode{root}
		layers [][]int
		layer  = 1
	)
	for len(nodes) > 0 {
		var (
			values    = make([]int, 0, len(nodes))
			nextNodes = make([]*TreeNode, 0)
		)

		for i := len(nodes) - 1; i >= 0; i-- {
			node := nodes[i]
			if node == nil {
				continue
			}
			values = append(values, node.Val)
			if layer%2 == 1 {
				// 从左到右遍历，先加左再加右
				nextNodes = append(nextNodes, node.Left, node.Right)
			} else {
				// 从右到左遍历，先加右再加左
				nextNodes = append(nextNodes, node.Right, node.Left)
			}
		}
		if len(values) > 0 {
			layers = append(layers, values)
		}
		nodes = nextNodes
		layer++
	}
	return layers
}
