from typing import List
import json

class Solution:
    def minSideJumps(self, obstacles: List[int]) -> int:
        dp = [1e8] * 3
        dp[1] = 0
        for ob in obstacles:
            if ob > 0:
                dp[ob-1] = 1e8
            mn = min(dp)
            for i, n in enumerate(dp):
                dp[i] = min(n, mn + 1)
            if ob > 0:
                dp[ob-1] = 1e8
        return min(dp)



s = Solution()
with open("/Users/it/Desktop/mininum.json") as f:
    d = json.load(f)
print(s.minSideJumps(d))
# print(s.minSideJumps([0,2,1,0,3,0]))