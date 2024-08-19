package main

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connectDFS(left, right *Node) {
	if left == nil {
		return
	}
	// 把两个点接起来
	left.Next = right
	connectDFS(left.Left, left.Right)
	connectDFS(left.Right, right.Left) // 左的右和右的左接起来
	connectDFS(right.Left, right.Right)
}

// 最直观的方法还是层序遍历
func connect(root *Node) *Node {
	if root == nil {
		return root
	}
	connectDFS(root.Left, root.Right)
	return root
}
