package main

import "fmt"

func canPlaceFlowers(flowerbed []int, n int) bool {
	count := 0
	i := 1
	for _, f := range flowerbed {
		if f == 0 {
			i++
		} else {
			count += (i - 1) / 2
			i = 0
		}
	}
	count += i / 2
	return count >= n
}

func main() {
	f := []int{1, 0, 0, 0, 1}
	fmt.Println(canPlaceFlowers(f, 2))
}
