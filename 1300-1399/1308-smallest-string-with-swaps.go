package main

import (
	"fmt"
	"sort"
)

func smallestStringWithSwaps(s string, pairs [][]int) string {
	var bytes = []byte(s)
	sort.Slice(pairs, func(i, j int) bool {
		p1, p2 := pairs[i], pairs[j]
		if bytes[p1[0]] > bytes[p2[0]] {
			return true
		} else if bytes[p1[0]] < bytes[p2[0]] {
			return false
		}
		return bytes[p1[1]] > bytes[p2[1]]
	})
	fmt.Println(pairs)
	for _, pair := range pairs {
		a, b := pair[0], pair[1]
		// if bytes[a] > bytes[b] {
		bytes[a], bytes[b] = bytes[b], bytes[a]
		// }
	}
	return string(bytes)
}

func main() {
	fmt.Println(smallestStringWithSwaps(
		"dcab",
		[][]int{
			{0, 3}, {1, 2}, {0, 2},
		}))
}
