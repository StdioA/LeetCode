package main

func longestCommonPrefix(strs []string) string {
	var (
		prefixLen int
		valid     = true
	)

	for i := 0; valid; i++ {
		var lastByte byte

		for _, str := range strs {
			if i >= len(str) {
				valid = false
				break
			}
			if lastByte == 0 {
				lastByte = str[i]
			} else if str[i] != lastByte {
				valid = false
				break
			}
		}
		if valid {
			prefixLen = i + 1
		}
	}
	return strs[0][:prefixLen]
}
