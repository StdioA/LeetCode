# coding: utf-8

class Solution:
    # @param n, an integer
    # @return an integer
    def climbStairs(self, n):
	    l = dict()					#dict的妙用（不用初始化一个长长的list了）
	    l[1] = 1
	    l[2] = 2
	    for x in range(3, n+1):
	    	l[x] = l[x-1] + l[x-2]
	    return l[n]

a = Solution()
for x in range(1, 10):
	print a.climbStairs(x)