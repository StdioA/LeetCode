# coding: utf-8

class Solution(object):
    def numIslands(self, grid):
        """
        :type grid: List[List[str]]
        :rtype: int
        """
        if not grid:
            return 0
            
        length, width = len(grid), len(grid[0])
        visited = [[False]*width for _ in range(length)]
        direction = ((1, 0), (-1, 0), (0, 1), (0, -1))
        total = 0

        def visit(x, y):
            visited[x][y] = True

            for dx, dy in direction:
                nx, ny = x+dx, y+dy
                if 0 <= nx < length and 0 <= ny < width:
                    if (not visited[nx][ny]) and grid[nx][ny]=="1":
                        visit(nx, ny)

        for i, line in enumerate(grid):
            for j, _ in enumerate(line):
                if not visited[i][j] and grid[i][j] == "1":
                    total += 1
                    visit(i, j)

        return total

a = Solution()
print a.numIslands([["1","0"], ["0","1"]])
