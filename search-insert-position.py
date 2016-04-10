# coding: utf-8

class Solution(object):
    def searchInsert(self, nums, target):
        """
        :type nums: List[int]
        :type target: int
        :rtype: int
        """
        try:
            i = nums.index(target)
        except ValueError:
            return len(filter(lambda x: x < target, nums))            
        else:
            return i

a = Solution()
print a.searchInsert([1,3,5,6], 0)
