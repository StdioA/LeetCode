# coding: utf-8

class Solution(object):
    def minPathSum(self, grid):
        """
        :type grid: List[List[int]]
        :rtype: int
        """
        m, n = len(grid), len(grid[0])
        dist = []

        for i in range(m):
            temp = []
            for j in range(n):
                temp.append(0)
            dist.append(temp)

        for i in range(m):
            for j in range(n):
                if i==0 and j==0:
                    dist[i][j] = grid[0][0]

                if i == 0:
                    dist[i][j] = dist[i][j-1]+grid[i][j]
                elif j == 0:
                    dist[i][j] = dist[i-1][j]+grid[i][j]
                else:
                    dist[i][j] = min(dist[i][j-1], dist[i-1][j])+grid[i][j]

        return dist[-1][-1]

obj = Solution()
grid = [[1,0,0,0,1],[2,2,2,2,1],[2,2,2,2,1]]
print(obj.minPathSum(grid))
