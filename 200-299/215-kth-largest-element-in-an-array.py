# coding: utf-8
from collections import defaultdict


class Solution:
    def findKthLargest(self, nums: List[int], k: int) -> int:
        bucket = defaultdict(int)
        for num in nums:
            bucket[num] += 1

        i = 10000
        while i >= -10000:
            if bucket[i] >= k:
                return i
            k -= bucket[i]
            i -= 1
        return 0
