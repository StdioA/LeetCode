package main

func separateDigits(nums []int) []int {
	var result []int
	for i := len(nums) - 1; i >= 0; i-- {
		num := nums[i]
		for num > 0 {
			result = append(result, num%10)
			num /= 10
		}
	}
	length := len(result)
	for i := 0; i < length/2; i++ {
		result[i], result[length-1-i] = result[length-1-i], result[i]
	}
	return result
}
