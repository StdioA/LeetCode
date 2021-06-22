package main

// 找到第一个下降的数字，再倒过来找到下降点，减一，后面的全都是 9
// 123321 - 122999
func monotoneIncreasingDigits(n int) int {
	var (
		digits []int
		num    = n
	)
	for num > 0 {
		digits = append(digits, num%10)
		num /= 10
	}
	var (
		l     = len(digits)
		index int
	)
	// 找到第一个下降的数字
	for index = l - 1; index >= 1; index-- {
		if digits[index-1] < digits[index] {
			break
		}
	}
	// 所有数字都是上升的，如 1234
	if index == 0 {
		return n
	}
	// 反过来找到第一个小于等于的数字
	for ; index < l-1; index++ {
		if digits[index+1] < digits[index] {
			break
		}
	}
	// 这一位减 1，后面全部变成 9
	// 123321 -> 12 2999
	digits[index]--
	for i := index - 1; i >= 0; i-- {
		digits[i] = 9
	}

	var (
		result int
		valid  bool
	)
	// 100 -> 099
	for i := len(digits) - 1; i >= 0; i-- {
		// Remove leading zeroes
		valid = valid || (digits[i] > 0)
		if valid {
			result = result*10 + digits[i]
		}
	}
	return result
}
