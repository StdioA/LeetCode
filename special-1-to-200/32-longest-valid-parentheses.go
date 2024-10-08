package main

func longestValidParentheses_bruteforce(s string) int {
	// 时间复杂度 O(n^2)
	// 构建左括号和右括号的前缀和，方便后续统计左右的个数
	// 不过这一步好像也不需要，因为后续遍历还会累加的
	var (
		lefts  = make([]int, len(s)+1)
		rights = make([]int, len(s)+1)
	)
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			lefts[i+1], rights[i+1] = lefts[i]+1, rights[i]
		} else {
			lefts[i+1], rights[i+1] = lefts[i], rights[i]+1
		}
	}
	var maxLen int
	for i := 1; i <= len(s); i++ {
		for j := i; j <= len(s); j++ {
			accLeft := lefts[j] - lefts[i-1]
			accRight := rights[j] - rights[i-1]
			if accRight > accLeft {
				// 如果右括号比左括号还多，那么以 s[i-1] 开头的字符串就已经非法了
				break
			}

			if accLeft > 0 && accLeft == accRight {
				if maxLen < 2*accLeft {
					maxLen = 2 * accLeft
				}
			}
		}
	}
	return maxLen
}

func longestValidParentheses(s string) int {
	type stackItem struct {
		index             int
		validLengthBefore int
	}
	// 记录每一个未被匹配的左括号的位置
	var (
		stack             = []stackItem{}
		maxLen            int
		validLengthBefore int
	)
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, stackItem{
				index:             i,
				validLengthBefore: validLengthBefore,
			})
			validLengthBefore = 0
		} else if len(stack) > 0 {
			// 遇到右括号，就要强制 pop 出来一个左括号，并累计有效长度
			lastItem := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			length := i - lastItem.index + 1 + lastItem.validLengthBefore
			// 判断大小
			if length > maxLen {
				maxLen = length
			}
			validLengthBefore = length
		} else {
			// 如果右括号太多，导致栈被排空，此时就需要重置累计长度
			validLengthBefore = 0
		}
	}
	return maxLen
}
