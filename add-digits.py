# coding: utf-8

class Solution(object):
    def addDigits(self, num):
        """
        :type num: int
        :rtype: int
        """
        nstr = str(num)
        while len(nstr)>1:
            n = sum([int(x) for x in nstr])
            nstr = str(n)
            
        return int(nstr)
