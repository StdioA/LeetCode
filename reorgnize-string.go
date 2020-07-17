package main

import (
	"fmt"
	"strings"
)

func reorganizeString(S string) string {
	charMap := make(map[byte]int)
	for _, char := range S {
		charMap[byte(char)]++
	}
	odd := byte(0)
	length := len(S)
	chars := make([]byte, length)
	index := 0
	for char, occurance := range charMap {
		if occurance%2 == 1 {
			if odd > 0 {
				return ""
			} else {
				odd = char
			}
		}
		for i := 1; i <= occurance/2; i++ {
			chars[index] = char
			chars[length-index-1] = char
			index++
		}
	}
	if odd > 0 {
		chars[index] = odd
	}
	var builder strings.Builder
	builder.Write(chars)
	return builder.String()
}

func main() {
	fmt.Println(reorganizeString("aab"))
	fmt.Println(reorganizeString("aabb"))
	fmt.Println(reorganizeString("aaab"))
}
