# coding: utf-8


class Solution(object):
    def moveZeroes(self, nums):
        """
        :type nums: List[int]
        :rtype: void Do not return anything, modify nums in-place instead.
        """
        index = 0
        length = len(nums)
        while index < length:
            if nums[index] == 0:
                nums.remove(0)
                nums.append(0)
                length -= 1
            else:
                index += 1
