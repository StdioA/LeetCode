func addBST(root *TreeNode, base int) int {
    if root == nil {
        return base
    }
    rightBase := addBST(root.Right, base)
    root.Val += rightBase
	leftVal := addBST(root.Left, root.Val)
    return leftVal
}

// Same as convert-bst-to-greater-tree
func bstToGst (root *TreeNode) *TreeNode {
    addBST(root, 0)
    return root  
}