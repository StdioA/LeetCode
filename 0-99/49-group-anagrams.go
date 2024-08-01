package main

import (
	"strconv"
	"strings"
)

func freqToStr(freqMap map[byte]int) string {
	var builder strings.Builder
	for i := byte('a'); i <= byte('z'); i++ {
		if freqMap[i] > 0 {
			builder.WriteByte(i)
			builder.WriteString(strconv.Itoa(freqMap[i]))
		}
	}
	return builder.String()
}

func groupAnagrams(strs []string) [][]string {
	// 用 [26]byte 更快
	var anagrams = make(map[string][]string)

	for _, str := range strs {
		freq := make(map[byte]int)
		for i := 0; i < len(str); i++ {
			freq[str[i]]++
		}
		freqStr := freqToStr(freq)
		anagrams[freqStr] = append(anagrams[freqStr], str)
	}
	var result = make([][]string, 0, len(anagrams))
	for _, value := range anagrams {
		result = append(result, value)
	}
	return result
}
