package main

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	var digits = make([]byte, 0, 10)
	for x > 0 {
		digits = append(digits, byte(x%10))
		x /= 10
	}
	length := len(digits) / 2
	for i := 0; i < length; i++ {
		if digits[i] != digits[len(digits)-1-i] {
			return false
		}
	}
	return true
}
