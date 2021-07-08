# coding: utf-8

class Solution:
    # @param {integer[]} nums
    # @param {integer} target
    # @return {integer[]}
    def twoSum(self, nums, target):
        # Time complicity: O(n)
        dict_ = {}
        for index, num in enumerate(nums, 1):
            i = dict_.get(target-num, None)
            if i:
                return [i, index]
            dict_[num] = index

        # Failed! Time complicity: O(n^2)

        # for index, num in enumerate(nums):                    
        #     tarnum = target - num
        #     # if tarnum in nums:
        #     try:
        #         return [index+1, nums.index(tarnum)+1]
        #     except ValueError:
        #         continue

a = Solution()
print a.twoSum([2,7,11,15],9)
