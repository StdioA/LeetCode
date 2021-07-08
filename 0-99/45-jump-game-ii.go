package main

func jump(nums []int) int {
	var (
		index int
		count int
	)

	for index < len(nums)-1 {
		count++
		maxJump := nums[index]

		// Can jump to the last index
		if index+maxJump >= len(nums)-1 {
			break
		}

		maxIndex := 0
		maxNextJump := 0
		for i := index + 1; i <= index+maxJump; i++ {
			if nums[i] == 0 {
				// Never jump into 0
				continue
			}
			nextJump := i + nums[i]
			if nextJump > maxNextJump {
				maxNextJump = nextJump
				maxIndex = i
			}
		}
		// Jump to that index
		index = maxIndex
	}
	return count
}
