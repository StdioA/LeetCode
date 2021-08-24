package main

import "sort"

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

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		ii, ij := intervals[i], intervals[j]
		if ii[0] < ij[0] {
			return true
		} else if ii[0] > ij[0] {
			return false
		}
		return ii[1] < ij[1]
	})

	var stack [][]int
	for _, interval := range intervals {
		for len(stack) > 0 {
			top := stack[len(stack)-1]
			if top[1] >= interval[0] {
				interval[0] = min(top[0], interval[0])
				interval[1] = max(top[1], interval[1])
				stack = stack[:len(stack)-1]
			} else {
				break
			}
		}
		stack = append(stack, interval)
	}
	return stack
}
