package main

import "fmt"

type DayT struct {
	day, temp int
}

func dailyTemperatures(T []int) []int {
    length := len(T)
    result := make([]int, length)
    stack := make([]DayT, 0)

    for i, temp := range T {
		j := len(stack) - 1
		for ; j >= 0; j-- {
			dayT := stack[j]
			if dayT.temp >= temp {
				break
			}
			result[dayT.day] = i - dayT.day
		}
		stack = stack[:j+1]
		stack = append(stack, DayT{i, temp})
	}
	return result
}


func main() {
	fmt.Println(dailyTemperatures([]int{73, 74, 75, 71, 69, 72, 76, 73}))
}
