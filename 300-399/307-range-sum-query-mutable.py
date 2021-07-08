# https://leetcode.com/problems/range-sum-query-mutable/
# https://blog.csdn.net/Small_Orange_glory/article/details/81290634
class NumArray:

    def __init__(self, nums: 'List[int]'):
        self.nums = nums
        self.bit = {}                       # 注意：树状数组的下标从 1 开始
        for i, num in enumerate(nums, 1):
            self.bit[i] = num               # 更新当前节点
            j = i - 2
            while j >= i - self.lowbit(i):  # 将下面的节点之和更新到当前节点
                self.bit[i] += nums[j]
                j -= 1

    def lowbit(self, num):
        return num & (-num)

    def update(self, i: 'int', val: 'int') -> 'None':
        delta = val - self.nums[i]
        self.nums[i] = val                  # 更新当前节点
        j = i + 1
        while j <= len(self.nums):          # 将当前节点的更新通知到当前节点的所有父节点
            self.bit[j] += delta
            j += self.lowbit(j)

    def sums(self, k):
        """
        求出树中 [1, k] 区间的和
        也就是 nums 数组中 [0, k - 1] 的和
        """
        ans = 0
        i = k
        while i > 0:                        # 向左上方依次加总所有节点的和
            ans += self.bit[i]
            i -= self.lowbit(i)             # 7(111) = 7(111) + 6(110) + 4(100)
        return ans

    def sumRange(self, i: 'int', j: 'int') -> 'int':
        return self.sums(j + 1) - self.sums(i)

a = NumArray([1, 3, 5])

print(a.sumRange(0, 2))     # 9
a.update(1, 2)
print(a.sumRange(0, 2))     # 8

