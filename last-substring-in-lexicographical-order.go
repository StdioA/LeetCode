package main

import (
	"fmt"
	"strings"
)

// 后缀数组：https://oi-wiki.org/string/sa/
// 退化到 O(n^2) 的算法，被用例卡傻了
func lastSubstring(s string) string {
	var (
		maxByte    = s[0]
		maxIndices = []int{0}
	)

	for i := 1; i < len(s); i++ {
		if s[i] > maxByte {
			maxByte = s[i]
			maxIndices = []int{i}
		} else if s[i] == maxByte {
			maxIndices = append(maxIndices, i)
		}

		lastIndex := maxIndices[len(maxIndices)-1]
		offset := i - lastIndex
		for {
			lastIndex := maxIndices[len(maxIndices)-1]
			if lastIndex+offset > i {
				break
			}
			var (
				firstIndex        = maxIndices[0]
				filteredIndices   = []int{firstIndex}
				maxByteWithOffset = s[firstIndex+offset]
			)
			for _, index := range maxIndices[1:] {
				if s[index+offset] > maxByteWithOffset {
					maxByteWithOffset = s[index+offset]
					filteredIndices = []int{index}
				} else if s[index+offset] == maxByteWithOffset {
					filteredIndices = append(filteredIndices, index)
				}
			}
			fmt.Println(i, maxIndices, filteredIndices)
			offset++
			maxIndices = filteredIndices
		}
	}

	// Find max string
	var maxString string
	for _, index := range maxIndices {
		subStr := s[index:]
		if strings.Compare(subStr, maxString) == 1 {
			maxString = subStr
		}
	}
	return maxString
}

func main() {
	// fmt.Println(lastSubstring("zzzazzxa"))
	fmt.Println(lastSubstring("abcdexyzfzz"))
}
