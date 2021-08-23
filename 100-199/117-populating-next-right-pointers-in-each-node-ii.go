package main

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	var layer = []*Node{root}
	for len(layer) > 0 {
		var nextLayer []*Node
		// Set next pointers
		for i := 0; i < len(layer); i++ {
			node := layer[i]
			if i > 0 {
				layer[i-1].Next = node
			}
			if node.Left != nil {
				nextLayer = append(nextLayer, node.Left)
			}
			if node.Right != nil {
				nextLayer = append(nextLayer, node.Right)
			}
		}
		layer = nextLayer
	}
	return root
}
