package main

func xorQueries(arr []int, queries [][]int) []int {
	var (
		preXor  = make([]int, len(arr))
		results = make([]int, len(queries))
	)
	preXor[0] = arr[0]
	for i := 1; i < len(arr); i++ {
		preXor[i] = preXor[i-1] ^ arr[i]
	}
	for i, query := range queries {
		result := preXor[query[1]]
		if query[0] > 0 {
			result ^= preXor[query[0]-1]
		}
		results[i] = result
	}
	return results
}
