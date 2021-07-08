package main

import "fmt"

func findTheDifference(s string, t string) byte {
	ds, dt := make(map[rune]int), make(map[rune]int)
	for _, rs := range s {
		ds[rs]++
	}
	for _, rt := range t {
		dt[rt]++
	}
	fmt.Println(ds, dt)
	for k, v := range dt {
		if ds[k] != v {
			return byte(k)
		}
	}
	return byte(0)
}

func main() {
	fmt.Println(findTheDifference("abcd", "dceba"))
}
