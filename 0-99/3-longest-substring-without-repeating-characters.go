package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	var (
		start, end = 0, 0
		max        int
		freq       = make(map[byte]bool)
	)
	for end < len(s) {
		// Start move forward
		for start <= end && freq[s[end]] {
			freq[s[start]] = false
			start++
		}
		// End move forward
		for end < len(s)-1 {
			freq[s[end]] = true
			if freq[s[end+1]] {
				break
			}
			end++
		}
		length := end - start + 1
		if length > max {
			max = length
		}
		end++
	}
	return max
}

func main() {
	fmt.Println(lengthOfLongestSubstring("bbbbb"))
}
