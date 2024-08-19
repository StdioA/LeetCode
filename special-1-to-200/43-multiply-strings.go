package main

import (
	"strings"
)

func multiply(a, b string) string {
	// 高精度乘法
	var (
		as = make([]byte, len(a))
		bs = make([]byte, len(b))
	)
	for i := 0; i < len(a); i++ {
		as[len(a)-i-1] = a[i] - '0'
	}
	for i := 0; i < len(b); i++ {
		bs[len(b)-i-1] = b[i] - '0'
	}
	result := make([]byte, len(a)+len(b))
	for i, an := range as {
		for j, bn := range bs {
			result[i+j] += an * bn
			if result[i+j] >= 10 {
				result[i+j+1] += result[i+j] / 10
				result[i+j] %= 10
			}
		}
	}
	var (
		buf   strings.Builder
		valid bool
	)
	buf.Grow(len(result))
	for i := len(result) - 1; i >= 0; i-- {
		n := result[i]
		if n > 0 {
			valid = true
		}
		if valid {
			buf.WriteByte(n + '0')
		}
	}
	s := buf.String()
	if s == "" {
		return "0"
	}
	return s
}
