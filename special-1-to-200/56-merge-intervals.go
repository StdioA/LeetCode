package main

import "sort"

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		ii, ij := intervals[i], intervals[j]
		if ii[0] == ij[0] {
			return ii[1] < ij[1]
		}
		return ii[0] < ij[0]
	})
	var result = [][]int{intervals[0]}
	for _, interval := range intervals {
		for i := len(result) - 1; i >= 0; i-- {
			// 从栈顶开始合并
			lastInterval := result[i]
			if interval[0] <= lastInterval[1] {
				// 合并之后，把该元素 pop 掉
				// interval 里是合并过的值，最后再放回栈里
				interval = []int{min(interval[0], lastInterval[0]), max(interval[1], lastInterval[1])}
				result = result[:i]
			} else {
				break
			}
		}
		result = append(result, interval)
	}
	return result
}
