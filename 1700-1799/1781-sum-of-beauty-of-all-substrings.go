package main

import "fmt"

func getMinMax(arr []int) (int, int) {
	var min, max int
	for _, n := range arr {
		if n > 0 && (min == 0 || n < min) {
			min = n
		}
		if n > max {
			max = n
		}
	}
	return min, max
}

func beautySum(s string) int {
	var (
		freqs = make([][]int, len(s))
		freq  = make([]int, 26)
	)
	// make prefix num
	for i := 0; i < len(s); i++ {
		c := s[i]
		freq[c-'a']++
		tmp := make([]int, 26)
		copy(tmp, freq)
		freqs[i] = tmp
	}
	var total int
	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s); j++ {
			// calculate beauty for s[i:j+1]
			freq := make([]int, 26)
			copy(freq, freqs[j])
			if i > 0 {
				for c := 0; c < 26; c++ {
					freq[c] -= freqs[i-1][c]
				}
			}
			minNum, maxNum := getMinMax(freq)
			fmt.Println(s[i:j+1], minNum, maxNum, freq)
			total += maxNum - minNum
		}
	}
	return total
}

func main() {
	println(beautySum("aabcbaa"))
}
