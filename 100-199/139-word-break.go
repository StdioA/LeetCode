package main

import "fmt"

func wordBreak(s string, wordDict []string) bool {
	result := make([]bool, len(s)+1)
	result[0] = true

	for i := 1; i <= len(s); i++ {
		for _, word := range wordDict {
			lw := len(word)
			if (i >= lw) && result[i-lw] && (s[i-lw:i] == word) {
				result[i] = true
				break
			}
		}
	}
	return result[len(s)]
}

func main() {
	s := "applepenapple"
	wordDict := []string{"apple", "pen"}
	fmt.Println(wordBreak(s, wordDict))
}
