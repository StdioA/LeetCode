package main

import "fmt"

func findAnagrams(s string, p string) []int {
	var target = make(map[byte]int)
	for i := 0; i < len(p); i++ {
		target[p[i]]++
	}

	var (
		start   = 0
		end     = 0
		current = map[byte]int{}
		result  []int
	)

	for end = 0; end < len(s); end++ {
		b := s[end]

		// If character not in p, then clear the map
		if target[b] == 0 {
			start = end + 1
			current = map[byte]int{}
			continue
		} else if current[b] >= target[b] {
			// Frequency exceeded, move start pointer forward
			for start <= end && current[b] >= target[b] {
				current[s[start]]--
				start++
			}
		}

		current[b]++
		// Check frequence equal
		var equal = true
		for k, v := range target {
			if current[k] != v {
				equal = false
				break
			}
		}

		if equal {
			result = append(result, start)
		}
	}
	return result
}

func main() {
	s := "cbaebabacd"
	p := "abc"
	fmt.Println(findAnagrams(s, p))
}
