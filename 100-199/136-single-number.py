# coding: utf-8

class Solution:
    # @param A, a list of integer
    # @return an integer
    def singleNumber(self, A):
        n = 0
        for num in A:
            n ^= num
        return n