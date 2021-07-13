package main

func combine(n int, k int) [][]int {
	if k <= 0 {
		return [][]int{
			{},
		}
	}
	if n <= k {
		res := make([]int, n)
		for i := 1; i <= n; i++ {
			res[i-1] = i
		}
		return [][]int{res}
	}

	var (
		results = combine(n-1, k)
		partial = combine(n-1, k-1)
	)
	for _, p := range partial {
		results = append(results, append(p, n))
	}
	return results
}
