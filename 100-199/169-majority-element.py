# coding: utf-8

class Solution:
    # @param num, a list of integers
    # @return an integer
    def majorityElement(self, num):
        dic = {}
        for n in num:
        	if dic.has_key(n):
        		dic[n] += 1
        	else:
        		dic[n] = 1
        l = len(num)
        for n, count in dic.items():
        	if count > l/2:
        		return n