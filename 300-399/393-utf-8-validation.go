package main

import "fmt"

func validUtf8(data []int) bool {
	var (
		index  int
		amount int
	)
	for index < len(data) {
		// Calculate how many "1"s in the leading byte
		amount = 0
		b := data[index]
		for mask := 7; mask > 0; mask-- {
			if b&(1<<mask) > 0 {
				amount++
			} else {
				break
			}
		}

		if amount == 0 {
			index++
			continue
		} else if amount == 1 || amount > 4 {
			return false
		}
		// validate follow bytes
		index++
		for amount--; amount > 0 && index < len(data); amount-- {
			// check whether byte is 10xxxxxx
			if data[index]&(0b11<<6) != 0b10<<6 {
				return false
			}
			// move to the next byte
			index++
		}
	}
	return amount == 0
}

func main() {
	fmt.Println(validUtf8([]int{250, 145, 145, 145, 145}))
}
