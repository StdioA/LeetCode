# coding: utf-8


class Solution(object):
    def singleNumber(self, nums):
        """
        :type nums: List[int]
        :rtype: List[int]
        """
        d = {}
        for num in nums:
            if not num in d:
                d[num] = True
            else:
                del d[num]
                
        return d.keys()
