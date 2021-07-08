# coding: utf-8
# UNFINISHED

class Solution(object):
    def productExceptSelf(self, nums):
        """
        :type nums: List[int]
        :rtype: List[int]
        """
        # https://leetcode.com/discuss/87865/how-from-o-n-to-o-1
        n= len(nums)
        result = [1]*n
        left = right = 1

        for i, num in enumerate(nums):
            result[i] *= left
            result[n-1-i] *= right;

            left *= nums[i];
            right *= nums[n-1-i]

        return result

    def productExceptSelf_RE(self, nums):
        """
        :type nums: List[int]
        :rtype: List[int]
        """
        import operator
        pro = reduce(operator.mul, nums)
        return [operator.div(pro, i) for i in nums]
        
a = Solution()
print a.productExceptSelf([1,2,3,4])
