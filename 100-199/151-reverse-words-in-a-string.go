package main

import "fmt"

func reverseWords(s string) string {
	var (
		bytes   = []byte(s)
		start   int
		current int
		inWords bool
	)

	// Flip the whole string
	length := len(bytes)
	for i := 0; i < length/2; i++ {
		bytes[i], bytes[length-1-i] = bytes[length-1-i], bytes[i]
	}

	for _, b := range bytes {
		if b != ' ' {
			if inWords {
				// Put character in the position
				bytes[current] = b
				current++
			} else {
				// Start counting the word
				inWords = true
				bytes[current] = b
				start = current
				current++
			}
		} else if inWords {
			// Reverse word between [start, current]
			length := current - start
			for j := 0; j < length/2; j++ {
				bytes[start+j], bytes[current-1-j] = bytes[current-1-j], bytes[start+j]
			}

			// Add space after the word
			bytes[current] = ' '
			current++
			inWords = false
		}
	}
	// Reverse the last word
	if inWords {
		length := current - start
		for j := 0; j < length/2; j++ {
			bytes[start+j], bytes[current-1-j] = bytes[current-1-j], bytes[start+j]
		}
		current++
	}
	return string(bytes[:current-1])
}

func main() {
	s := "hello      world"
	fmt.Printf(`"%s"`, reverseWords(s))
}
