package main

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func intervalIntersection(firstList [][]int, secondList [][]int) [][]int {
	var (
		firstIndex, secondIndex int
		result                  [][]int
	)
	for firstIndex < len(firstList) && secondIndex < len(secondList) {
		fseg, sseg := firstList[firstIndex], secondList[secondIndex]
		prev, next := max(fseg[0], sseg[0]), min(fseg[1], sseg[1])
		if prev <= next {
			result = append(result, []int{prev, next})
		}

		if fseg[1] < sseg[1] {
			firstIndex++
		} else if fseg[1] > sseg[1] {
			secondIndex++
		} else {
			firstIndex++
			secondIndex++
		}
	}
	return result
}
