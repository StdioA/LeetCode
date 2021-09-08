package main

func dfs(res *[]string, s string, open, close_, max int) {
	if len(s) == max*2 {
		*res = append(*res, s)
	}
	if open < max {
		dfs(res, s+"(", open+1, close_, max)
	}
	if close_ < open {
		dfs(res, s+")", open, close_+1, max)
	}
}

func generateParenthesis(n int) []string {
	var res = []string{}
	dfs(&res, "", 0, 0, n)
	return res
}
