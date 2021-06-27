from typing import List

class Solution:
    def rotateGrid(self, grid: List[List[int]], k: int) -> List[List[int]]:
        m, n = len(grid), len(grid[0])
        rounds = min(m, n) // 2

        directions = ((1, 0), (0, 1), (-1, 0), (0, -1), (0, 0))
        for r in range(rounds):
            x, y = r, r
            d = 0

            numbers = []
            while d < 4:
                numbers.append(grid[x][y])
                nx, ny = x + directions[d][0], y + directions[d][1]
                # 判断越界
                if not ((0 <= nx + (r * directions[d][0]) < m) and (0 <= ny + (r * directions[d][1]) < n)):
                    d += 1
                    nx, ny = x + directions[d][0], y + directions[d][1]
                x, y = nx, ny
                if (x, y) == (r, r):
                    break

            roundmod = k % len(numbers)
            numbers = numbers[-roundmod:] + numbers[:-roundmod]

            # 把数放回去
            x, y = r, r
            d = 0
            while d < 4:
                grid[x][y] = numbers.pop(0)
                nx, ny = x + directions[d][0], y + directions[d][1]
                # 判断越界
                if not ((0 <= nx + (r * directions[d][0]) < m) and (0 <= ny + (r * directions[d][1]) < n)):
                    d += 1
                    nx, ny = x + directions[d][0], y + directions[d][1]

                x, y = nx, ny
                if (x, y) == (r, r):
                    break
        return grid

a = Solution()
print(a.rotateGrid([[1,2,3,4],[5,6,7,8],[9,10,11,12],[13,14,15,16]], 2))
