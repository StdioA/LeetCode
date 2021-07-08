package main

import "strings"

func simplifyPath(path string) string {
	var (
		segments = strings.Split(path, "/")
		stack    []string
	)
	for _, seg := range segments {
		switch seg {
		case "", ".":
			continue
		case "..":
			l := len(stack)
			if l > 0 {
				stack = stack[:l-1]
			}
		default:
			stack = append(stack, seg)
		}
	}
	return "/" + strings.Join(stack, "/")
}
