# coding: utf-8

class Solution(object):
    def canJump(self, nums):
        """
        :type nums: List[int]
        :rtype: bool
        """
        maxr = 0
        for i, n in enumerate(nums):
            if i > maxr:
                return False
            maxr = max(maxr, i+n)
        return True


a = Solution()
print a.canJump([3,2,1,0,4])
