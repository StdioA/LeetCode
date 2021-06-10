package main

import "fmt"

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 离谱的 O(n^2) dp
func maxAreaSlow(height []int) int {
	var dp = make([]int, len(height))
	for i := range height {
		for j := i + 1; j < len(height); j++ {
			dp[j] = max(dp[j], min(height[i], height[j])*(j-i))
		}
	}
	ans := dp[0]
	for _, val := range dp {
		ans = max(ans, val)
	}
	return ans
}

// 双指针
func maxArea(height []int) int {
	var (
		start, end = 0, len(height) - 1
		size       = min(height[start], height[end]) * (end - start)
	)
	for start < end {
		s, e := start, end
		if height[start] < height[end] {
			for start < len(height) && height[start] <= height[s] {
				start++
			}
		} else {
			for end > 0 && height[end] <= height[e] {
				end--
			}
		}
		size = max(size, min(height[start], height[end])*(end-start))
	}
	return size
}

func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}
