package main

import "fmt"

func longestAwesome(s string) int {
	var (
		posMap    = make(map[int]int) // key: bitmap, val: minimum index
		bitMap    int
		maxLength int
		masks     = make([]int, 11)
	)
	posMap[0] = -1
	for i := 0; i < 10; i++ {
		masks[i] = 1 << i
	}
	fmt.Println(masks)

	for i := 0; i < len(s); i++ {
		bitMap ^= 1 << (s[i] - '0')
		if _, ok := posMap[bitMap]; !ok {
			posMap[bitMap] = i
		}
		for _, mask := range masks {
			targetMap := bitMap ^ mask
			if index, ok := posMap[targetMap]; ok && (i-index) > maxLength {
				maxLength = i - index
			}
		}
	}
	return maxLength
}

func main() {
	fmt.Println(longestAwesome("213123"))
}
