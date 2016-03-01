# coding: utf-8


class Solution(object):
    def singleNumber(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        d = {}
        for num in nums:
            if not num in d:
                d[num] = 1
            elif d[num] == 1:
                d[num] = 2
            else:
                del d[num]
                
        return d.keys()[0]
