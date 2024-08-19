package main

func generateParenthesis(n int) []string {
	return generate(n, 0, 0)
}

func concat(prefix string, strings []string) []string {
	var result = make([]string, len(strings))
	for i, str := range strings {
		result[i] = prefix + str
	}
	return result
}

func generate(n, open, close int) []string {
	// 递归生成，open 为左括号数量，close 为右括号数量
	// 由于里面有太多的字符串拷贝，所以实际效率和空间占用都比较高
	if close == n {
		return []string{""}
	}
	var result []string
	if open > close {
		partial_result := concat(")", generate(n, open, close+1))
		result = append(result, partial_result...)
	}
	if open < n {
		partial_result := concat("(", generate(n, open+1, close))
		result = append(result, partial_result...)
	}
	return result
}
