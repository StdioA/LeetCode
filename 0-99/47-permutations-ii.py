# coding: utf-8

class Solution(object):
    def permuteUnique(self, nums):
        """
        :type nums: List[int]
        :rtype: List[List[int]]
        """
        import itertools
        l = list(itertools.permutations(nums))
        l = list(set(l))
        return [list(o) for o in l]
