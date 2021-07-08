from typing import List

class Solution:
    def countServers(self, grid: List[List[int]]) -> int:
        row_sum = list(map(sum, grid))
        col_sum = list(map(sum, zip(*grid)))
        return sum(
            c > 0 and (row_sum[i] > 1 or col_sum[j] > 1)
            for i, row in enumerate(grid)
            for j, c in enumerate(row)
        )


a = Solution()
print(a.countServers([[1,1,0,0],[0,0,1,0],[0,0,1,0],[0,0,0,1]]))