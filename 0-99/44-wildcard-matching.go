package main

import "fmt"

// 补丁太多了还是没过
func isMatch(s string, p string) bool {
	var (
		lenS           = len(s)
		lenP           = len(p)
		sIndex, pIndex int
	)
	for sIndex < lenS && pIndex < lenP {
		fmt.Println(sIndex, pIndex)
		curByte := p[pIndex]
		switch curByte {
		case '?':
			sIndex++
			pIndex++
		case '*':
			for pIndex < len(p) && p[pIndex] == '*' {
				pIndex++
			}
			if pIndex == lenP {
				return true
			}

			// TODO: 后缀和优化
			asterisks := 0
			for pp := pIndex; pp < len(p); pp++ {
				if p[pp] == '*' {
					asterisks++
				}
			}

			// Match s from last to current character
			nextSIndex := (lenS - (lenP - pIndex - asterisks))
			// fmt.Println("PP", pIndex, asterisks, nextSIndex)
			// 这段匹配算法是错的，应该匹配后面所有不是 * 的部分
			for nextSIndex >= sIndex {
				// fmt.Println("match", nextSIndex, pIndex, s[nextSIndex], p[pIndex])
				if s[nextSIndex] == p[pIndex] || p[pIndex] == '?' {
					break
				}
				nextSIndex--
			}
			// fmt.Println("PP Done", pIndex, nextSIndex)
			if nextSIndex > sIndex {
				sIndex = nextSIndex
			}
		default:
			if s[sIndex] != p[pIndex] {
				return false
			}
			sIndex++
			pIndex++
		}
	}
	// s should be fully matched
	if sIndex != lenS {
		return false
	}
	// Check if p has trailing alterisks
	for pIndex < lenP {
		if p[pIndex] != '*' {
			return false
		}
		pIndex++
	}
	return true
}

func main() {
	// fmt.Println(isMatch("adceb", "*a*b"))               // true
	// fmt.Println(isMatch("acdcb", "a*c?b"))              // false
	// fmt.Println(isMatch("mississippi", "m??*ss*?i*pi")) // false
	// fmt.Println(isMatch("aaaa", "***a"))                // true
	// fmt.Println(isMatch("b", "*?*?"))                   // false
	// fmt.Println(isMatch("aa", "*"))                     // true
	fmt.Println(isMatch("mississippi", "m*iss*")) // true
	// mis si ss  ippi
	// m??  * ss* ?i*pi
}
