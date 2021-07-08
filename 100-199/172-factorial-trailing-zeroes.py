# coding: utf-8

class Solution:
    # @return an integer
    def trailingZeroes(self, n):
    	if n/5:
    		return n/5 + self.trailingZeroes(n/5)
    	else:
    		return 0

a = Solution()
print a.trailingZeroes(100)