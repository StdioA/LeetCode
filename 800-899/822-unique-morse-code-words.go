package main

import "fmt"

func uniqueMorseRepresentations(words []string) int {
	charset := []string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}
	set := make(map[string]bool)
	for _, word := range words {
		morse := ""
		for _, char := range word {
			morse += charset[int(char)-int('a')]
		}
		set[morse] = true
	}
	return len(set)
}

func main() {
	s := []string{"gin", "zen", "gig", "msg"}
	fmt.Println(uniqueMorseRepresentations(s))
}
