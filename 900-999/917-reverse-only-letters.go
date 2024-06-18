package main

func isLetter(b byte) bool {
	return 'a' <= b && b <= 'z' || 'A' <= b && b <= 'Z'
}

func reverseOnlyLetters(s string) string {
	var (
		i, j  = 0, len(s) - 1
		bytes = []byte(s)
	)
	for i < j {
		for i < len(bytes) && !isLetter(bytes[i]) {
			i++
		}
		for j >= 0 && !isLetter(bytes[j]) {
			j--
		}
		if i < j {
			bytes[i], bytes[j] = bytes[j], bytes[i]
			i++
			j--
		}
	}
	return string(bytes)
}
