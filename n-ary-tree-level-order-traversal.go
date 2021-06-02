/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func levelOrder(root *Node) [][]int {
	if root == nil {
		return nil
	}
	var (
		current = []*Node{root}
		result  [][]int
	)
	for len(current) > 0 {
		var (
			level []int
			next  []*Node
		)
		for _, node := range current {
			level = append(level, node.Val)
			next = append(next, node.Children...)
		}
		result = append(result, level)
		current = next
	}
	return result
}