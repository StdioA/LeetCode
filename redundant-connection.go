package main

import "fmt"

func findRedundantConnection(edges [][]int) []int {
	pointSet := make(map[int]int)

	find := func (n int) int {
		for pointSet[n] > 0 {
			n = pointSet[n]
		}
		return n
	}

	for _, edge := range edges {
		start, end := edge[0], edge[1]

		s, e := find(start), find(end)
		if s == e {
			return edge
		} else {
			pointSet[s] = e
		}
	}
	return []int{}
}

func main() {
	edges := [][]int{
		{1, 4},
		{3, 4},
		{1, 3},
		{1, 2},
		{4, 5},
	}
	fmt.Println(findRedundantConnection(edges))
}
