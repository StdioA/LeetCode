func missingNumber(nums []int) int {
	var (
		l = len(nums)
		s = l * (l + 1) / 2
	)
	for _, n := range nums {
		s -= n
	}
	return s
}