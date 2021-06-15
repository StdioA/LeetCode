package main

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func insert(intervals [][]int, newInterval []int) [][]int {
	var (
		curInterval = newInterval
		result      [][]int
	)
	for i, interval := range intervals {
		iStart, iEnd := interval[0], interval[1]
		cStart, cEnd := curInterval[0], curInterval[1]

		switch {
		case iEnd < cStart:
			result = append(result, interval)
		case cEnd < iStart:
			result = append(result, curInterval)
			return append(result, intervals[i:]...)
		default:
			curInterval[0] = min(iStart, cStart)
			curInterval[1] = max(iEnd, cEnd)
		}
	}
	return append(result, curInterval)
}
