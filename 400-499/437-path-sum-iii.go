package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 还是直接用前缀和来做
func dfs(node *TreeNode, targetSum int) (map[int]int, int) {
	if node == nil {
		return map[int]int{}, 0
	}
	// Left & Right
	var val = node.Val
	lpSum, lTotal := dfs(node.Left, targetSum)
	rpSum, rTotal := dfs(node.Right, targetSum)

	// 更新前缀和，把所有前缀和都加上当前值
	pSumCounter := make(map[int]int)
	for k, v := range lpSum {
		pSumCounter[k+val] += v
	}
	for k, v := range rpSum {
		pSumCounter[k+val] += v
	}
	pSumCounter[val]++

	// Total 为路径起始点低于等于（？）当前节点的路径条数
	total := lTotal + rTotal + pSumCounter[targetSum]
	return pSumCounter, total
}

func pathSum(root *TreeNode, targetSum int) int {
	_, total := dfs(root, targetSum)
	return total
}

// 把前缀和计数器的方向倒过来，可以省去 copy counter 的大量时间
func pathSumFast(root *TreeNode, targetSum int) int {
	var (
		cache = map[int]int{0: 1}
		res   int
	)
	pathSumUtil(root, targetSum, 0, cache, &res)
	return res
}

func pathSumUtil(root *TreeNode, target int, currSum int, cache map[int]int, res *int) {
	if root == nil {
		return
	}
	currSum += root.Val
	*res += cache[currSum-target]
	cache[currSum]++
	pathSumUtil(root.Left, target, currSum, cache, res)
	pathSumUtil(root.Right, target, currSum, cache, res)
	cache[currSum]--
}
