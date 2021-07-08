# coding: utf-8

# 算格雷码

class Solution:
	# @return a list of integers
	def grayCode(self, n):
		numl = range(2**(n))					#题目说的太不清楚了吧…
		ans = []
		for x in numl:
			lb = [int(x) for x in bin(x)[2:]]
			tb = ''
			tb = str(lb[0])
			for t in range(1, len(lb)):
				tb += str(lb[t]^lb[t-1])		#算法是乱搞的，正规一点应该用位运算(好麻烦)，实在懒得鼓捣了
			ans.append(int(tb, base=2))
		return ans

a = Solution()
print a.grayCode(2)