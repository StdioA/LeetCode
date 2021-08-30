package main

func minArray(numbers []int) int {
	var i, j = 0, len(numbers) - 1
	for i < j {
		mid := (i + j) / 2
		switch {
		case numbers[mid] > numbers[j]:
			i = mid + 1
		case numbers[mid] < numbers[j]:
			j = mid
		default:
			j--
		}
	}
	return numbers[i]
}
