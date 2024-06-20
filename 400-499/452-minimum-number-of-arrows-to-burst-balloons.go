package main

import "sort"

func findMinArrowShots(points [][]int) int {
	// 不断取交集
	var segments [][]int
	sort.Slice(points, func(i, j int) bool {
		pi, pj := points[i], points[j]
		if pi[0] == pj[0] {
			return pi[1] < pj[1]
		}
		return pi[0] < pj[0]
	})

	for _, point := range points {
		if len(segments) == 0 {
			segments = append(segments, point)
			continue
		}
		lastSeg := segments[len(segments)-1]
		if lastSeg[1] >= point[0] {
			// 两个区间有交叉，取交集
			if point[0] > lastSeg[0] {
				lastSeg[0] = point[0]
			}
			if point[1] < lastSeg[1] {
				lastSeg[1] = point[1]
			}
		} else {
			segments = append(segments, point)
		}
	}
	return len(segments)
}
