# coding: utf-8

#找一坨字符串的最长公共前缀
#减少比较次数
class Solution:
	# @return a string
	def longestCommonPrefix(self, strs):
		if len(strs) == 0:
			return ''
		longest = reduce(min, map(len, strs))
		if longest == 0:
			return ''

		for i in range(len(strs)-1):
			for x in range(longest):
				if strs[i][x] != strs[i+1][x]:
					longest = x
					break

		return strs[0][:longest]

a = Solution()
s = ['a', 'b']
print a.longestCommonPrefix(s)