# coding: utf-8

class Solution:
    # @param s, a string
    # @return an integer
    def titleToNumber(self, s):
    	n = [ord(x)-64 for x in list(s)][::-1]
    	l = len(n)
    	num = 0
    	for x in range(l):
    		num += n[x]*(26**x)
    	return num


a = Solution()
print a.titleToNumber('AB')