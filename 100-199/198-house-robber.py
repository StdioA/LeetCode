# coding: utf-8

class Solution:
    # @param {integer[]} nums
    # @return {integer}
    def rob(self, nums):
        l = len(nums)

        if l == 0:
            return 0
        elif l in [1,2]:
            return max(nums)

        nums[2] += nums[0]
        for i in range(3, l):
            nums[i] += max(nums[i-2],nums[i-3])

        return max(nums[-1], nums[-2])