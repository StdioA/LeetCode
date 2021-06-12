package main

import "sort"

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)

	var (
		count  int
		sIndex int
		gIndex int
	)

	for sIndex < len(s) && gIndex < len(g) {
		if s[sIndex] >= g[gIndex] {
			count++
			gIndex++
		}
		sIndex++
	}
	return count
}
