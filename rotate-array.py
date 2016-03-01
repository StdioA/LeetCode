# coding: utf-8

class Solution:
    # @param nums, a list of integer
    # @param k, num of steps
    # @return nothing, please modify the nums list in-place.
    def rotate(self, nums, k):
        for x in range(k):
            num = nums[-1]
            nums.pop()
            nums.insert(0, num)
        return None

a = Solution()
l = [1,2,3,4,5,6,7]
a.rotate(l, 3)
print l