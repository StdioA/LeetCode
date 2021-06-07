package main

func genList(node *TreeNode) (head, tail *TreeNode) {
	head = node
	tail = node

	left, right := node.Left, node.Right
	node.Left = nil
	if left != nil {
		lhead, ltail := genList(left)
		head = lhead
		ltail.Right = node
		tail = node
	}
	if right != nil {
		rhead, rtail := genList(right)
		tail.Right = rhead
		tail = rtail
	}
	return head, tail
}

func increasingBST(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	head, _ := genList(root)
	return head
}
