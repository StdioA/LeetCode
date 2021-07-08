# coding: utf-8

#算法纯属乱搞T^T

class Solution:
	# @return an integer
	def atoi(self, s):
		s = s.strip()
		begin, end = 0, len(s)
		if s == '':
			return 0

		haveSign = False
		if s[0] in '+-':
			begin = 1
			haveSign = True								#是否有符号(+1, -2)

		for index in range(begin,len(s)):
			if not ('0' <= s[index] <= '9'):
				end = index
				break

		if begin == end:
			return 0
		elif not ('0' <= s[end-1] <= '9'):
			return 0


		if haveSign:
		    begin -=1

		num = int(s[begin:end])
		if num > 2147483647:
			return 2147483647
		elif num < -2147483648:
			return -2147483648
		return num


a = Solution()
print a.atoi("1")
