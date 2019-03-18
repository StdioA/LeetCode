package main

import "fmt"

func diStringMatch(S string) []int {
	l, r, arr := 0, len(S), make([]int, len(S))
	for i, c := range S {
		if c == 'I' {
			arr[i] = l
			l++
		} else if c == 'D' {
			arr[i] = r
			r--
		}
	}
	return append(arr, r)
}

func main() {
	fmt.Println(diStringMatch("IDID"))
}
