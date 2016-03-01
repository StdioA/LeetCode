# coding: utf-8

class Solution:
    # @param s, a string
    # @return an integer
    def lengthOfLastWord(self, s):
        l = s.strip().split(' ')
        return len(l[-1])