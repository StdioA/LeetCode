package main

import (
	"encoding/json"
	"fmt"
)

func isBipartite(graph [][]int) bool {
	var (
		nodes  = make([]int, len(graph))
		search func(i int) bool
	)
	search = func(i int) bool {
		if i >= len(graph) {
			return true
		}
		var (
			line      = graph[i]
			nodeTypes = []int{1, 2}
		)
		if nodes[i] > 0 {
			nodeTypes = []int{nodes[i]}
		}
		for _, nodeType := range nodeTypes {
			var (
				available   = true
				anotherType = 3 - nodeType
				backup      = make(map[int]int)
			)
			nodes[i] = nodeType
			for _, n := range line {
				if nodes[n] == nodeType {
					available = false
				}
				backup[n] = nodes[n]
				nodes[n] = anotherType
			}
			if available {
				if res := search(i + 1); res {
					return true
				}
			}
			for i, n := range backup {
				nodes[i] = n
			}
		}
		return false
	}
	nodes[0] = 1
	return search(0)
}

func main() {
	var (
		input    [][]int
		inputStr = []byte("[[1],[0,3],[3],[1,2]]")
	)

	json.Unmarshal(inputStr, &input)
	res := isBipartite(input)
	fmt.Println(res)
}
