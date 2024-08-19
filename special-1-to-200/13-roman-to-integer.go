package main

func romanToInt(s string) int {
	var symbolMap = map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	var result int
	for i := 0; i < len(s); i++ {
		b := s[i]
		switch {
		case i+1 < len(s) && symbolMap[b] < symbolMap[s[i+1]]:
			result -= symbolMap[b]
		default:
			result += symbolMap[b]
		}
	}
	return result
}
