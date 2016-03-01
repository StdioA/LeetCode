# coding: utf-8

class Solution:
    # @return a list of lists of integers
    def generate(self, numRows):
    	if numRows == 0:
    		return []
    	lnum = [[1]]
        for n in range(2, numRows+1):
        	cur = lnum[-1]
        	next = [1]*n
        	for k in range(1, n-1):
        		next[k] = cur[k-1]+cur[k]
        	lnum.append(next)
        return lnum

a = Solution()
print a.generate(5)