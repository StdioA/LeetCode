package main

import "strings"

func intToRoman(num int) string {
	// 打表模拟
	var table = [][]string{
		{"I", "V", "X"},
		{"X", "L", "C"},
		{"C", "D", "M"},
		{"M"},
	}
	var digits = make([]int, 0, 4)
	for num > 0 {
		digits = append(digits, num%10)
		num /= 10
	}
	var builder strings.Builder
	for i := len(digits) - 1; i >= 0; i-- {
		row := table[i]

		switch digit := digits[i]; digit {
		case 0, 1, 2, 3:
			for x := 1; x <= digit; x++ {
				builder.WriteString(row[0])
			}
		case 4:
			builder.WriteString(row[0])
			builder.WriteString(row[1])
		case 5, 6, 7, 8:
			builder.WriteString(row[1])
			for x := 6; x <= digit; x++ {
				builder.WriteString(row[0])
			}
		case 9:
			builder.WriteString(row[0])
			builder.WriteString(row[2])
		}
	}
	return builder.String()
}
