package main

import (
	"strings"
)

func convert(s string, numRows int) string {
	// 第一行和最后一行单独处理
	// 中间行用同一种处理方法
	if numRows == 1 {
		return s
	}
	var (
		builder strings.Builder
		cycle   = 2 * (numRows - 1)
	)
	// 第一行
	for i := 0; i < len(s); i += cycle {
		builder.WriteByte(s[i])
	}
	// 中间行
	for row := 1; row < numRows-1; row++ {
		for p := 0; p < len(s); p += cycle {
			// 写入当前组的第 i 和第 n-i 个字符
			p1, p2 := p+row, p+(cycle-row)
			if p1 < len(s) {
				builder.WriteByte(s[p1])
			}
			if p2 < len(s) {
				builder.WriteByte(s[p2])
			}
		}
	}
	// 最后一行
	for i := numRows - 1; i < len(s); i += cycle {
		builder.WriteByte(s[i])
	}
	return builder.String()
}
