package main

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

func maxArea(height []int) int {
	// 不断收缩的窗口
	var (
		start, end = 0, len(height) - 1
		size       = min(height[start], height[end]) * (end - start)
	)
	for start < end {
		s, e := start, end
		// 宽度最大的情况下，如果要找到更大的面积，那么高度一定要比当前高度更高
		// 所以是一个不断去找更高高度的过程
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
