package main

func countTriplets(arr []int) int {
	var (
		prefixSum    int
		prefixXorMap = make(map[int][]int)
		total        int
	)

	prefixXorMap[0] = append(prefixXorMap[0], -1)
	for i := 0; i < len(arr); i++ {
		prefixSum = prefixSum ^ arr[i]
		prefixXorMap[prefixSum] = append(prefixXorMap[prefixSum], i)
	}

	for _, indices := range prefixXorMap {
		length := len(indices)
		for i := 0; i < length-1; i++ {
			for j := i + 1; j < length; j++ {
				total += indices[j] - indices[i] - 1
			}
		}
	}
	return total
}

// 假设 prefix[n] = arr[0:n] 的异或和，0 <= n <= len(arr)
// 如果存在 i, j, k, 则 a ^ b == 0，arr[i:k+1] 的异或和为 0，可推出 prefix[i] == prefix[k+1]
// 所以求出 prefix[n] 每个取值对应的坐标即可
// i, k 确定后，任意满足 i < j <= k 的 j 值均能满足 a == b，即 (i, k) 对应的三元组个数为 k - i
// 注意一个边缘 case 是 i == 0 的情况，计算 prefix 前要加一个 prefix[0] 进去
