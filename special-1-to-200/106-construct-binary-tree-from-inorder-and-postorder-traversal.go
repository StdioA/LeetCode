package main

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}

	var (
		rootValue = postorder[len(postorder)-1]
		root      = &TreeNode{Val: rootValue}
	)
	// fetch index from inorder
	var index int
	for inorder[index] != rootValue {
		index++
	}
	root.Left = buildTree(inorder[:index], postorder[:index])
	root.Right = buildTree(inorder[index+1:], postorder[index:len(inorder)-1])
	return root
}
