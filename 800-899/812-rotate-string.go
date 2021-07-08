package main

import (
	"fmt"
	"strings"
)

func rotateString(A string, B string) bool {
	return len(A) == len(B) && strings.Contains(A+A, B)
}

func main() {
	fmt.Println(rotateString("abcde", "cdeab"))
}
