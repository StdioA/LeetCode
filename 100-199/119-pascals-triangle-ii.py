# coding: utf-8

class Solution:
    # @return a list of lists of integers
    def generate(self, rowIndex):
    	cur = [1]
        for n in range(rowIndex+2):
        	next = [1]*n
        	for k in range(1, n-1):
        		next[k] = cur[k-1]+cur[k]
        	cur = next
        return cur

a = Solution()
print a.generate(1)