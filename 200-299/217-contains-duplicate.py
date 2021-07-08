# coding: utf-8

class Solution:
    # @param {integer[]} nums
    # @return {boolean}
    def containsDuplicate(self, nums):
        dul = {}
        for num in nums:
            if dul.get(num, None):
                return True
            else:
                dul[num] = True
        return False

print Solution().containsDuplicate([1,2,3,4,5,3])