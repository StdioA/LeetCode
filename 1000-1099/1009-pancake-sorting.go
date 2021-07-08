package main

func reverse(slice []int) {
	var l = len(slice)
	for i := 0; i < l/2; i++ {
		slice[i], slice[l-1-i] = slice[l-1-i], slice[i]
	}
}

func pancakeSort(arr []int) []int {
	var (
		result []int
		total  = len(arr)
	)
	for i := total; i >= 1; i-- {
		// i is in place
		if arr[i-1] == i {
			continue
		}
		var index int
		for index = 0; index < i && arr[index] != i; index++ {
		}
		if index > 0 {
			result = append(result, index+1)
			reverse(arr[:index+1])
		}
		result = append(result, i)
		reverse(arr[:i])
	}
	return result
}

// 3 2 4 1
// Swap 4 to the first position, and then to the corresponding position
// Then move 3
// 3 - 4 2 3 1
// 4 - 1 3 2 4 - 4 done
// 2 - 3 1 2 4
// 3 - 2 1 3 4 - 3 done
// 1 - 1 2 3 4 - 2 done, thus 1 done
