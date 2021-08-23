from typing import List


class Solution:
    def longestConsecutive(self, nums: List[int]) -> int:
        ext = {}
        for num in nums:
            ext[num] = True

        maxLength = 0
        for i, val in ext.items():
            if not val:
                continue

            ext[i] = False
            upper = i + 1
            while ext.get(upper):
                ext[upper] = False
                upper += 1
            lower = i - 1
            while ext.get(lower):
                ext[lower] = False
                lower -= 1
            maxLength = max(maxLength, upper - lower - 1)

        return maxLength


a = Solution()
print(a.longestConsecutive([100, 4, 200, 1, 3, 2]))
