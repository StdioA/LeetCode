/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func dfs(root, p, q *TreeNode) (foundP, foundQ bool, node *TreeNode) {
	if root == nil {
		return
	}

	if root == p {
		foundP = true
	} else if root == q {
		foundQ = true
	}

	nodes := []*TreeNode{root.Left, root.Right}
	for _, n := range nodes {
		fp, fq, found := dfs(n, p, q)
		if found != nil {
			node = found
			return
		}
		foundP = foundP || fp
		foundQ = foundQ || fq
		if foundP && foundQ {
			node = root
			return
		}
	}
	return
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	_, _, node := dfs(root, p, q)
	return node
}
