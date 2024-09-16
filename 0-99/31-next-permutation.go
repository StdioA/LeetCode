package main

func reverse(nums []int) {
	for i := len(nums)/2 - 1; i >= 0; i-- {
		opp := len(nums) - 1 - i
		nums[i], nums[opp] = nums[opp], nums[i]
	}
}

func nextPermutation(nums []int) {
	length := len(nums)
	left, right := 0, length-1
	// 找到第一个升序的数对，左边的那个数就是需要交换的
	// 此时右边的所有数都是降序的
	for right > 0 && nums[right] <= nums[right-1] {
		right--
	}
	if right > 0 {
		left = right - 1
		right = length - 1
		// 找到第一个比左边的数还要大的数，并将其交换
		for right > left && nums[right] <= nums[left] {
			right--
		}
		if right > left {
			nums[left], nums[right] = nums[right], nums[left]
		}
		left++
	}
	// 交换之后，就要重排左边之后的数，使其变成升序
	reverse(nums[left:])
}


1 2 3 4
1 2 4 3
1 3 2 4
1 3 4 2
1 4 2 3
1 4 3 2
2 1 3 4

2 4 3 1
3 1 2 4


// func main() {
// 	var l []int
// 	l = []int{1, 2}
// 	nextPermutation(l)
// 	fmt.Println(l)
// 	l = []int{1, 3, 2}
// 	nextPermutation(l)
// 	fmt.Println(l)
// 	l = []int{1, 2, 3}
// 	nextPermutation(l)
// 	fmt.Println(l)
// }

// (l) / 2
// 3 -> i < 1
// 0 - 2
// 4 -> i < 2
// 0 - 3
// 5 -> i < 2
// 0 - 4
