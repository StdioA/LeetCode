package main

import "strings"

func wordPattern(pattern string, s string) bool {
	var (
		patterns = make(map[byte]string)
		segments = strings.Split(s, " ")
	)
	if len(pattern) != len(segments) {
		return false
	}
	for i := 0; i < len(pattern); i++ {
		p, word := pattern[i], segments[i]

		original, ok := patterns[p]
		if !ok {
			patterns[p] = word
		} else if original != word {
			return false
		}
	}
	var dup = make(map[string]struct{})
	for _, word := range patterns {
		_, ok := dup[word]
		if ok {
			return false
		}
		dup[word] = struct{}{}
	}
	return true
}
