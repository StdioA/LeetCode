# coding: utf-8

class Solution:
    # @param n, an integer
    # @return an integer
    def hammingWeight(self, n):
        binnum = bin(n)
        length = len([x for x in binnum if x == '1'])
        return length

a = Solution()
print a.hammingWeight(11)