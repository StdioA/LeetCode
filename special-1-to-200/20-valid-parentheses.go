package main

func isValid(s string) bool {
	var (
		stack      []byte
		bracketMap = map[byte]byte{
			')': '(',
			']': '[',
			'}': '{',
		}
	)
	for i := 0; i < len(s); i++ {
		current := s[i]
		oppose, ok := bracketMap[current]
		if ok {
			// 检查栈非空
			if len(stack) == 0 {
				return false
			}
			// 检查栈顶元素
			if stack[len(stack)-1] != oppose {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, current)
		}
	}
	return len(stack) == 0
}
