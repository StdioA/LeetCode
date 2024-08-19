package main

func combinationSum(candidates []int, target int) [][]int {
	if target == 0 {
		return [][]int{{}}
	}

	var results [][]int
	for index, candidate := range candidates {
		var cand_repeat []int
		for i := 1; i*candidate <= target; i++ {
			cand_repeat = append(cand_repeat, candidate)
			lastSum := combinationSum(candidates[index+1:], target-i*candidate)
			for _, sum := range lastSum {
				results = append(results, append(sum, cand_repeat...))
			}
		}
	}
	return results
}

// func main() {
// 	fmt.Println(combinationSum([]int{2, 5, 8, 4}, 10))
// }
