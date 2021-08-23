package main

import "fmt"

func findRepeatedDnaSequences(s string) []string {
	var (
		bits = map[byte]int{
			'A': 0,
			'T': 1,
			'C': 2,
			'G': 3,
		}
		occursion = make(map[int]int)
		window    int
		mask      = 1<<20 - 1
		results   []string
	)
	if len(s) <= 10 {
		return nil
	}
	for i := 0; i < 9; i++ {
		window = window<<2 | bits[s[i]]
	}
	for i := 9; i < len(s); i++ {
		window = (window<<2 | bits[s[i]]) & mask
		fmt.Println(window)
		occursion[window]++
		if occursion[window] == 2 {
			results = append(results, s[i-9:i+1])
		}
	}
	return results
}

func main() {
	fmt.Println(findRepeatedDnaSequences("AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"))
}
