package main

func longestPalindrome(s string) int {
	var freqMap = make(map[byte]int)
	for i := 0; i < len(s); i++ {
		freqMap[s[i]]++
	}

	var (
		total int
		odd   int
	)
	for _, freq := range freqMap {
		if freq%2 == 0 {
			total += freq
		} else {
			odd = 1
			total += freq - 1
		}
	}
	return total + odd
}
