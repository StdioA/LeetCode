# coding: utf-8

class Solution:
    # @return an integer
    def reverse(self, x):
        if x < 0:
            negative = True
            x = -x
        else:
            negative = False
        k = int(str(x)[::-1])
        if k > 2147483647:
            return 0
        if negative:
            return -k
        else:
            return k