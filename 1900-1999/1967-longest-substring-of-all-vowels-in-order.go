package main

import "fmt"

func longestBeautifulSubstring(word string) int {
	var (
		vowels     = []byte("aeiou")
		state  int = -1
		count  int
		max    int
	)
	for i := 0; i < len(word); {
		c := word[i]
		switch {
		case state == -1 && c != vowels[0]:
			i++
		case state >= 0 && c == vowels[state]:
			count++
			i++
		case state < 4 && c == vowels[state+1]:
			state++
			count++
			i++
		default:
			if state == 4 && count > max {
				max = count
			}
			count = 0
			state = -1
		}
	}
	if state == 4 && count > max {
		max = count
	}
	return max
}

func main() {
	fmt.Println(longestBeautifulSubstring("aeiaaeiouuaei"))
}
