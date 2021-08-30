package main

func findNthDigit(n int) int {
	if n <= 9 {
		return n
	}

	var (
		bit   = 1
		digit = 1
	)
	// 从 1 开始，跳过 9, 90, 900 直至找到合适的位数
	for n > bit*digit*9 {
		n -= bit * digit * 9
		digit *= 10
		bit++
	}
	n--
	num := digit + (n / bit)
	for i := 0; i < bit-1-(n%bit); i++ {
		num /= 10
	}
	return num % 10
}
