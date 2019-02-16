class Solution:
    def threeSum(self, nums):
        """
        :type nums: List[int]
        :rtype: List[List[int]]
        """
        from collections import Counter
        nums = list(sorted(nums))
        counter = Counter(nums)
        
        def check(a, b, c):
            count = counter.copy()
            for i in (a, b, c):
                count[i] -= 1
                if count[i] < 0:
                    return False
            return True

        results = []
        length = len(nums)
        for i in range(length):
            for j in range(i + 1, length):
                a, b = nums[i], nums[j]
                c = -a - b
                if check(a, b, c):
                    results.append(tuple(sorted([a, b, -a - b])))
        results = list(set(results))
        results = list(map(list, results))
        return results


s = Solution()
print(s.threeSum([-1, 0, 1, 2, -1, 4]))
