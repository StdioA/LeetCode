package main

import (
	"bytes"
	"fmt"
)

func getPermutation(n int, k int) string {
	var (
		bucket = make([]bool, n+1)
		fact   = 1
		buffer bytes.Buffer
	)
	for i := 1; i <= n; i++ {
		bucket[i] = true
		fact *= i
	}

	for i := n; i > 0; i-- {
		fact /= i
		seq := (k - 1) / fact
		k -= seq * fact

		// Append seq-th number in the buffer
		for j := 1; j <= n; j++ {
			if bucket[j] {
				seq--
			}
			if seq == -1 {
				bucket[j] = false
				buffer.WriteByte(byte(j) + '0')
				break
			}
		}
	}
	return buffer.String()
}

func main() {
	for i := 1; i < 24; i++ {
		fmt.Println(i, getPermutation(4, i))
	}
}

// 1234
// 1243
// 1324
// 1342
// 1423
// 1432
// ----
// 2134
// 2143
// 2314

// 8: 2 -> 1 -> 2 -> 1
// (8-1) / 6 = 1 ... 2
// (2-1) / 2 = 0 ... 2
// (2-1) / 1 = 1 ... 0
// (1-1) / 1 = 0 ... 0

// 9: 2 -> 2 -> 1 -> 1
// (9-1) / 6 = 1 ... 3
// (3-1) / 2 = 1 ... 1
// (1-1) / 1 = 0 ... 1
// (1-1) / 1 = 0 ... 1
