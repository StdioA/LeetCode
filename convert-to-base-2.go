package main

import (
	"fmt"
	"strings"
)

// 真·位运算，写的很鬼扯，但是 A 了
func baseNeg2Biiiiiiiiitwise(n int) string {
	if n == 0 {
		return "0"
	}

	var (
		bits  = make([]int, 0, 21)
		carry bool
	)
	for n > 0 {
		bits = append(bits, n&1)
		n = n >> 1
	}
	// pad 0 to generate odd bits
	if len(bits)%2 == 0 {
		bits = append(bits, 0)
	}
	// pad two zeroes for carry
	bits = append(bits, 0, 0)

	for i := 1; i < len(bits); i += 2 {
		// Notice that bits are reversed, bits[i] is low bit
		num := bits[i] + bits[i+1]<<1
		if !carry {
			switch num {
			case 0, 2:
				continue
			case 1:
				bits[i+1], bits[i] = 1, 1
			case 3:
				bits[i+1], bits[i] = 0, 1
				carry = true
			}
		} else {
			carry = false
			switch num {
			case 0:
				bits[i+1], bits[i] = 1, 1
			case 1:
				bits[i+1], bits[i] = 1, 0
			case 2:
				carry = true
				bits[i+1], bits[i] = 0, 1
			case 3:
				carry = true
				bits[i+1], bits[i] = 0, 0
			}
		}
	}
	var (
		s     strings.Builder
		valid bool
	)
	for i := len(bits) - 1; i >= 0; i-- {
		if bits[i] > 0 {
			// Drop leading zeroes
			valid = true
		}
		if valid {
			if bits[i] == 1 {
				s.WriteRune('1')
			} else {
				s.WriteRune('0')
			}
		}
	}
	return s.String()
}

func baseNeg2(n int) string {
	var (
		lowBit  = n & 1
		num     = 1                 // Set lower bit to 1 for bit padding, otherwise the trailing zero will lost
		carry   bool                // Carry flag
		revBits = []int{0, 2, 1, 3} // Used to reverse bits in a word, e.g. 00 -> 00, 10 -> 01
	)
	n >>= 1
	for n > 0 {
		word := n & 3
		if !carry {
			carry = word == 3
			word = 3 - (word+3)%4 // 这里其实是打表… [0, 3, 2, 1]
		} else {
			carry = word >= 2
			word = 3 - word // 同上 [3, 2, 1, 0]
		}
		num = num<<2 + revBits[word]
		n >>= 2
	}
	if carry {
		num = num<<2 + 3
	}

	var builder strings.Builder
	for num > 1 {
		builder.WriteByte('0' + byte(num&1))
		num >>= 1
	}
	builder.WriteByte('0' + byte(lowBit))
	return builder.String()
}

func main() {
	fmt.Println(baseNeg2((31)))
	fmt.Println(baseNeg2Biiiiiiiiitwise((31)))
}

// 1 -2 4 -8 16 -32 64
// 1 1 0 0 0 1 0

// bin(x)

// 如果偶数位为 1，且高位为 0，则更高位置 1（进位为 1）
// 如果偶数为 1，且高位为 1，则进 2 位，将后面的两个高位置 1

// 2 + 4 = 6 = 16 - 8 - 2
// 2 + 4 + 8 = 14 = 16 - 2
// 2 + 4 + 16 = 22 = 64 - 32 - 8 - 2
// 2 + 4 + 8 + 16 = 30 = 64 - 32 - 2

// 进 0 位的情况：
// 0 0 -> 0 0
// 1 0 -> 1 0

// 进 1 位的情况：
// 0 1 -> 1 0

// 进 2 位的情况：
// 1 1 -> 0 1 +2

// +2 后
// 0 0 变成 1 1
// 0 1 变成 1 0
// 1 0 变成 0 1 + 2
// 1 1 变成 0 0 + 2
