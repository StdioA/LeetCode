package main

import (
	"fmt"
	"strings"
	"strconv"
)

func decodeString(s string) string {
	var (
		result, top string
		number int
	)
	stack := make([]string, 0, len(s))
	
	pop := func (stack []string) ([]string, string) {
		result := stack[len(stack)-1]
		return stack[:len(stack)-1], result
	}
	manipulate := func(base, operator string) string {
		// "a", "2" -> "aa"
		// "a", "b" -> "ba"
		num, err := strconv.Atoi(operator)
		if err != nil {
			base = operator + base
		} else {
			base = strings.Repeat(base, num)
		}
		return base
	}

    for _, r := range s {
		if r != ']' {
			if r >= '0' && r <= '9' {
				// 合并数字
				number = number * 10 + int(r - '0')
			} else {
				if number > 0 {
					stack = append(stack, strconv.Itoa(number))
					number = 0
				}
				stack = append(stack, string(r))
			}
		} else {
			stop := false
			result = ""
			for {
				stack, top = pop(stack)
				if top == "[" {
					// 遇到 [ 后再进行一次 pop 操作就退出
					stop = true
					continue
				}
				result = manipulate(result, top)
				if stop {
					break
				}
			}
			stack = append(stack, result)
		}
	}
	result = ""
	for len(stack) > 0 {
		stack, top = pop(stack)
		result = manipulate(result, top)
	}
    return result
}

func main() {
	fmt.Println(decodeString("a15[b]"))
}
