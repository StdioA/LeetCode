package main

func subsets(nums []int) [][]int {
	var (
		length = len(nums)
		result = make([][]int, 0, 1<<length)
	)
	// 递归也可以，不过位运算更快，内存拷贝更少
	for i := 0; i < 1<<length; i++ {
		var (
			t   = i
			j   int
			ans = []int{}
		)
		for t > 0 {
			if t&1 > 0 {
				ans = append(ans, nums[j])
			}
			t >>= 1
			j++
		}
		result = append(result, ans)
	}
	return result
}
