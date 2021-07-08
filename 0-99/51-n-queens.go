package main

import (
	"fmt"
	"strings"
)

func solveNQueens(n int) [][]string {
	var (
		y      = make([]bool, n)
		m      = make([]bool, n*2)
		s      = make([]bool, n*2)
		status = make([]int, n)
		result [][]string
		solve  func(layer, n int) // How to make self-call in an anonymous function?
	)

	sol := func(layer, n int) {
		if layer == n {
			var res = make([]string, 0, n)
			for i := 0; i < n; i++ {
				var (
					b     strings.Builder
					place = status[i]
				)
				for j := 0; j < n; j++ {
					if j == place {
						b.WriteRune('Q')
					} else {
						b.WriteRune('.')
					}
				}
				res = append(res, b.String())
			}
			result = append(result, res)
			return
		}

		for i := 0; i < n; i++ {
			if !(y[i] || m[i+layer] || s[i-layer+n]) {
				y[i] = true
				m[i+layer] = true
				s[i-layer+n] = true
				status[layer] = i
				solve(layer+1, n)
				y[i] = false
				m[i+layer] = false
				s[i-layer+n] = false
			}
		}
	}
	solve = sol
	solve(0, n)
	return result
}

func main() {
	fmt.Println(solveNQueens(4))
}
