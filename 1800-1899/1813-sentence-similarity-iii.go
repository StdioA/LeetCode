package main

import (
	"strings"
)

func areSentencesSimilar(sentence1 string, sentence2 string) bool {
	var (
		words1 = strings.Split(sentence1, " ")
		words2 = strings.Split(sentence2, " ")
	)
	if len(words1) == len(words2) {
		return sentence1 == sentence2
	}
	if len(words2) == 0 {
		return true
	}
	if len(words1) < len(words2) {
		words1, words2 = words2, words1
	}
	var startIndex, endIndex int
	for ; startIndex < len(words2); startIndex++ {
		if words1[startIndex] != words2[startIndex] {
			break
		}
	}
	for ; startIndex+endIndex < len(words2) && endIndex < len(words2); endIndex++ {
		if words1[len(words1)-1-endIndex] != words2[len(words2)-1-endIndex] {
			break
		}
	}
	// fmt.Println(startIndex, endIndex)
	if startIndex+endIndex == len(words2) {
		return true
	}
	return false
}
