# coding: utf-8

from math import log

class Solution:
	# @return a string
	def convertToTitle(self, num):
		s = ''
		while num:
			num -= 1
			s += chr(num%26 + 65)
			num /= 26
		return s[::-1]

a = Solution()
print a.convertToTitle(701)