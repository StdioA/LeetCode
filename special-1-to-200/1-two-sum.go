package main

func twoSum(nums []int, target int) []int {
	// 比较简单的模拟
	// 用 map 缓存 target - num，随后在遍历时，直接查看该数是否存在即可
	var sumMap = make(map[int]int)
	for i, num := range nums {
		if anotherIndex, ok := sumMap[num]; ok {
			return []int{anotherIndex, i}
		}
		sumMap[target-num] = i
	}
	return nil
}
