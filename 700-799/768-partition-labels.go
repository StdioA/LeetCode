package main

import (
	"sort"
)

type charIndex struct {
	first, last int
}

func partitionLabels(s string) []int {
	// 统计每个字符第一次和最后一次出现的坐标
	var indices = make([]charIndex, 26)
	for i := 0; i < len(s); i++ {
		// 1-based
		index, b := i+1, s[i]
		indiceIndex := int(b - 'a')
		if indices[indiceIndex].first == 0 {
			indices[indiceIndex].first = index
		}
		indices[indiceIndex].last = index
	}

	// 合并区间
	sort.Slice(indices, func(i, j int) bool {
		if indices[i].first == indices[j].first {
			return indices[i].last < indices[j].last
		}
		return indices[i].first < indices[j].first
	})
	var (
		lastIndex charIndex
		ans       []int
	)
	for _, index := range indices {
		// 去除无效元素
		if index.first == 0 {
			continue
		}

		if lastIndex.first == 0 {
			lastIndex = index
			continue
		}
		if index.first < lastIndex.last {
			if index.last > lastIndex.last {
				lastIndex.last = index.last
			}
		} else {
			// 可以分割了
			ans = append(ans, lastIndex.last-lastIndex.first+1)
			lastIndex = index
		}
	}
	if lastIndex.first > 0 {
		ans = append(ans, lastIndex.last-lastIndex.first+1)
	}

	return ans
}
