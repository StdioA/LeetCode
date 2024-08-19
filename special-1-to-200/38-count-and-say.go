package main

import (
	"fmt"
	"strconv"
	"strings"
)

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}
	lastResult := countAndSay(n - 1)
	var (
		builder strings.Builder
		count   = 1
		last    = lastResult[0]
	)
	for i := 1; i < len(lastResult); i++ {
		if lastResult[i] != last {
			builder.WriteString(strconv.Itoa(count))
			builder.WriteByte(last)
			count, last = 1, lastResult[i]
		} else {
			count++
		}
	}
	builder.WriteString(strconv.Itoa(count))
	builder.WriteByte(last)
	fmt.Println(n, builder.String())
	return builder.String()
}

// func main() {
// 	fmt.Println(countAndSay(5))
// }
