package main

func dfs(left, right *Node) {
	if left == nil {
		return
	}
	left.Next = right
	dfs(left.Left, left.Right)
	dfs(left.Right, right.Left)
	dfs(right.Left, right.Right)
}

// 非要搞常数空间复杂度，层序遍历它不香吗
func connect(root *Node) *Node {
	if root == nil {
		return root
	}
	dfs(root.Left, root.Right)
	return root
}
