package main

func twoSum(numbers []int, target int) []int {
	var s, e = 0, len(numbers) - 1
	for {
		sum := numbers[s] + numbers[e]
		switch {
		case sum == target:
			return []int{s + 1, e + 1}
		case sum < target:
			s++
		case sum > target:
			e--
		}
	}
}
