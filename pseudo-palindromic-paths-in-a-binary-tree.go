package main

func dfs(node *TreeNode, freq []int) int {
	freq[node.Val-1]++
	defer func() {
		freq[node.Val-1]--
	}()
	if node.Left == nil && node.Right == nil {
		// Check if palindrone
		var oddAmount int
		for _, val := range freq {
			oddAmount += val % 2
		}
		if oddAmount <= 1 {
			return 1
		}
		return 0
	}

	var total int
	if node.Left != nil {
		total += dfs(node.Left, freq)
	}
	if node.Right != nil {
		total += dfs(node.Right, freq)
	}
	return total
}

func pseudoPalindromicPaths(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var freq = make([]int, 9)
	return dfs(root, freq)
}
