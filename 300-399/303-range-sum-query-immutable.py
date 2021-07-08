# https://leetcode.com/problems/range-sum-query-immutable/
# 前缀和
class NumArray:

    def __init__(self, nums: 'List[int]'):
        self.sums = {-1: 0}
        for i, n in enumerate(nums):
            self.sums[i] = self.sums[i - 1] + n

    def sumRange(self, i: 'int', j: 'int') -> 'int':
        return self.sums[j] - self.sums[i - 1]
        
n = NumArray([-2, 0, 3, -5, 2, -1])
print(n.sumRange(2, 5))
