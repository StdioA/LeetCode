package main

import "sort"

func hIndex(citations []int) int {
	sort.Sort(sort.Reverse(sort.IntSlice(citations)))
	var hIndex int
	for i, citation := range citations {
		if citation >= i+1 {
			hIndex = i + 1
		} else {
			break
		}
	}

	return hIndex
}
