package main

func reverseWords(s string) string {
	var (
		bytes      = []byte(s)
		start, end int
	)
	for end < len(bytes) {
		// Find the end of the word
		for end < len(bytes) && bytes[end] != ' ' {
			end++
		}
		// Swap the word bytes[start:end]
		for i := 0; i < (end-start)/2; i++ {
			bytes[start+i], bytes[end-1-i] = bytes[end-1-i], bytes[start+i]
		}
		// Find next word
		for start = end; start < len(bytes) && bytes[start] == ' '; start++ {
		}
		end = start
	}
	return string(bytes)
}
