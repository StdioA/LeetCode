# coding: utf-8


class Solution:
    # @param a, a string
    # @param b, a string
    # @return a string
    def addBinary(self, a, b):
        na = int(a, base=2)
        nb = int(b, base=2)
        s = na+nb
        return bin(s)[2:]