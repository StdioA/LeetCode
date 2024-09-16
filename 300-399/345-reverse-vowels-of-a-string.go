package main

func reverseVowels(s string) string {
	var (
		bytes        = []byte(s)
		vowelIndices []int
	)
	for i, b := range bytes {
		switch b {
		case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
			vowelIndices = append(vowelIndices, i)
		}
	}
	for i := 0; i < len(vowelIndices)/2; i++ {
		indexa, indexb := vowelIndices[i], vowelIndices[len(vowelIndices)-i-1]
		bytes[indexa], bytes[indexb] = bytes[indexb], bytes[indexa]
	}
	return string(bytes)
}
