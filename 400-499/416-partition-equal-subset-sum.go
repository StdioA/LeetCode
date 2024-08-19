package main

func canPartition(nums []int) bool {
	var (
		sum   int
		count = make(map[int]int)
	)
	for _, num := range nums {
		sum += num
		count[num]++
	}
	if sum%2 != 0 {
		return false
	}
	return canPartitionDFS(nums, sum/2, make([]map[int]bool, len(nums)+1))
}

func canPartitionDFS(nums []int, target int, mem []map[int]bool) bool {
	if target == 0 {
		return true
	} else if target < 0 {
		return false
	}
	if len(nums) == 0 {
		return false
	}

	// 检查之前是否计算过这个结果，因为后续有可能还会查到它
	// 比如 [1, 2, 2, 3, 10] 就有可能会在取 [1, 3] 和取 [2, 2] 的同时去搜索 ([10], 5)
	lenMem := mem[len(nums)]
	if lenMem == nil {
		lenMem = make(map[int]bool)
		mem[len(nums)] = lenMem
	}
	ans, ok := lenMem[target]
	if ok {
		return ans
	}
	ans = canPartitionDFS(nums[1:], target, mem) || canPartitionDFS(nums[1:], target-nums[0], mem)
	lenMem[target] = ans
	return ans
}
