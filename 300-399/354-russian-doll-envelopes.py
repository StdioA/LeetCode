from typing import List
from bisect import bisect_left

class Solution:
    def maxEnvelopes(self, envelopes: List[List[int]]) -> int:
        envelopes.sort()
        
        result = 0
        dp = [0] * len(envelopes)
        # O(n^2) LIS
        for i, largerEnv in enumerate(envelopes):
            dp[i] = 1
            for j, env in enumerate(envelopes[:i]):
                if largerEnv[0] > env[0] and largerEnv[1] > env[1]:
                    dp[i] = max(dp[i], dp[j] + 1)
            result = max(result, dp[i])
        return result      

    # https://leetcode.com/problems/russian-doll-envelopes/discuss/1134011/JS-Python-Java-C%2B%2B-or-Easy-LIS-Solution-w-Explanation  
    def maxEnvelopes(self, E: List[List[int]]) -> int:
        E.sort(key=lambda x: (x[0], -x[1]))
        dp = []
        for _,height in E:
            left = bisect_left(dp, height)
            if left == len(dp): dp.append(height)
            else: dp[left] = height
        return len(dp)
        