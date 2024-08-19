package main

func flatten(root *TreeNode) {
	flatten2(root)
}

func flatten2(root *TreeNode) *TreeNode {
	// 返回值是链表的尾节点
	if root == nil {
		return nil
	}
	var left, right = root.Left, root.Right
	// newRight 应该是整个链表的最右节点，所以需要不断更新
	newRight := root
	if left != nil {
		// 处理左子树
		lRight := flatten2(left)
		// 左节点置空
		root.Left = nil
		// 把 root -> [left -> lRight] -> right 接起来
		root.Right = left
		lRight.Right = right
		newRight = lRight
	}
	if right != nil {
		rRight := flatten2(right)
		newRight = rRight
	}
	return newRight
}
