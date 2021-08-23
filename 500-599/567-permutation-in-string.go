package main

func check(c1, c2 []int) bool {
	for i, c := range c1 {
		if c != c2[i] {
			return false
		}
	}
	return true
}

func checkInclusion(s1 string, s2 string) bool {
	var (
		freq1 = make([]int, 26)
		freq2 = make([]int, 26)
	)
	for i := 0; i < len(s1); i++ {
		freq1[int(s1[i]-'a')]++
	}
	var i, j int
	for j < len(s2) {
		c2 := int(s2[j] - 'a')
		if freq1[c2] == 0 {
			// Restart counting
			freq2 = make([]int, 26)
			j++
			i = j
			continue
		}
		freq2[c2]++
		if freq1[c2] == freq2[c2] && check(freq1, freq2) {
			return true
		}
		// freq exceeded, slide i forward
		for i < j && freq2[c2] > freq1[c2] {
			freq2[int(s2[i]-'a')]--
			i++
		}
		j++
	}
	return false
}
