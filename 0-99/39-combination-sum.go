func combinationSum(candidates []int, target int) [][]int {
	res := make([][]int, 0)
	for i, candidate := range candidates {
		if candidate > target {
			continue
		} else if candidate == target {
			res = append(res, []int{candidate})
		} else {
			partRes := combinationSum(candidates[i:], target-candidate)
			for _, r := range partRes {
				res = append(res, append([]int{candidate}, r...))
			}
		}
	}
	return res
}