package main

// func min(a, b int) int {
// 	if a < b {
// 		return a
// 	}
// 	return b
// }

// func max(a, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }

func maximumGap(nums []int) int {
	// 桶排序的变体，基础为容斥原理
	// 假设有 n 个数，其范围为 [a, b]，那么我们将 [a, b] 区间均分为 n+1 个桶，并将所有数放入对应的桶内
	// 这种情况下，至少会有一个桶是空的，而这个桶的两侧就可能存在最大区间
	// 把所有空桶找到，并计算左侧区间的最大数和右侧区间的最小数之差，它们的最大值就是整个数列的最大差
	const (
		upperLimit = int(1e9 + 1)
		lowerLimit = int(-1)
	)
	var minNum, maxNum = upperLimit, lowerLimit
	for _, num := range nums {
		minNum, maxNum = min(minNum, num), max(maxNum, num)
	}
	if maxNum == minNum {
		return 0
	}
	// 构建 n+1 个 bucket，gap 应为 (max-min)/n
	var gap = (maxNum - minNum) / len(nums)
	// 为了避免 0 gap （如 [1,1,3,3]），gap 最小为 1
	// 因此 bucket 的数量也要动态调整
	if gap == 0 {
		gap = 1
	}
	var (
		bucketAmount = (maxNum-minNum)/gap + 1
		mins         = make([]int, bucketAmount)
		maxs         = make([]int, bucketAmount)
	)
	for i := range mins {
		mins[i], maxs[i] = upperLimit, lowerLimit
	}
	// 维护每个区间的最小值和最大值
	for _, num := range nums {
		bucketIndex := int((num - minNum) / gap)
		mins[bucketIndex] = min(mins[bucketIndex], num)
		maxs[bucketIndex] = max(maxs[bucketIndex], num)
	}
	// 用双指针检查相邻区间的最小和最大值，取他们的差
	var (
		i      = 0 // 第一个桶一定不会是空的，但最后一个桶有可能会
		maxGap = -1
	)
	for i < len(mins) {
		var j int
		// 找到下一个非空桶
		for j = i + 1; j < len(mins) && mins[j] == upperLimit; j++ {
		}
		if j < len(mins) {
			// 计算两区间之差，并更新最大值
			maxGap = max(maxGap, mins[j]-maxs[i])
		}
		i = j
	}
	return maxGap
}
