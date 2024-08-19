package main

func jump(nums []int) int {
	// 目标：尽可能跳的更远
	// 不需要 dp，只需要尽可能找到能跳到更远的格子就行了
	if len(nums) == 1 {
		return 0
	}
	var (
		current int
		jump    int
	)

	for current < len(nums) {
		num := nums[current]
		// 可以直接跳到，跳过去就结束了
		if current+num >= len(nums)-1 {
			jump++
			break
		}
		// 选一个可以跳的更远的点
		var nextJump, nextJumpIndex int
		for i := current + 1; i <= current+num; i++ {
			if nums[i]+num+i > nextJump {
				nextJump = nums[i] + num + i
				nextJumpIndex = i
			}
		}
		jump++
		current = nextJumpIndex
	}
	return jump
}

// func main() {
// 	// fmt.Println(jump([]int{2, 3, 0, 1, 4}))
// 	fmt.Println(jump([]int{1, 1, 1, 1, 1}))
// }
