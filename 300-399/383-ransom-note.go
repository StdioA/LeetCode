package main

func canConstruct(ransomNote string, magazine string) bool {
	var freqM = make(map[byte]int)
	for i := 0; i < len(magazine); i++ {
		freqM[magazine[i]]++
	}
	for i := 0; i < len(ransomNote); i++ {
		b := ransomNote[i]
		freqM[b]--
		if freqM[b] < 0 {
			return false
		}
	}
	return true
}
