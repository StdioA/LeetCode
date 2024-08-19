package main

import (
	"fmt"
	"strings"
)

func products(segments []string) []string {
	// 生成笛卡尔积
	if len(segments) == 0 {
		return []string{""}
	}
	seg := segments[0]
	previous := products(segments[1:])
	var results = make([]string, 0, len(seg)*len(previous))
	for i := 0; i < len(seg); i++ {
		for _, prevSeg := range previous {
			var builder strings.Builder
			builder.Grow(1 + len(prevSeg))
			builder.WriteByte(seg[i])
			builder.WriteString(prevSeg)
			results = append(results, builder.String())
		}

	}
	return results
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	var (
		strMap = map[byte]string{
			'2': "abc",
			'3': "def",
			'4': "ghi",
			'5': "jkl",
			'6': "mno",
			'7': "pqrs",
			'8': "tuv",
			'9': "wxyz",
		}
		segments = make([]string, len(digits))
	)
	for i := 0; i < len(digits); i++ {
		segments[i] = strMap[digits[i]]
	}
	return products(segments)
}

func main() {
	fmt.Println(letterCombinations("23"))
}
