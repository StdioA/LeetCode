func searchBST(root *TreeNode, val int) *TreeNode {
	for {
		if root == nil {
			return root
		}
		if val == root.Val {
			return root
		}
		if val < root.Val {
			root = root.Left
		} else if val > root.Val {
			root = root.Right
		}
	}
}
