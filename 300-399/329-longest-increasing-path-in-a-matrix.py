# coding: utf-8

class Solution(object):
    def longestIncreasingPath(self, matrix):
        """
        :type matrix: List[List[int]]
        :rtype: int
        """
        if not matrix:
            return 0

        ans = 0
        length = len(matrix)
        width = len(matrix[0])
        dir_vec = [(0, 1), (0, -1), (1, 0), (-1, 0)]
        dp = [[0]*width for i in range(length)]                                 # 记忆数组，避免重复dfs
        
        def dfs(x, y):
            ans = 1
            for dx, dy in dir_vec:
                # 判断越界
                if x+dx<0 or x+dx>=length or y+dy<0 or y+dy>=width:
                    continue
                elif matrix[x+dx][y+dy]>matrix[x][y]:
                    if dp[x+dx][y+dy]:
                        ans = max(ans, dp[x+dx][y+dy]+1)
                    else:
                        dp[x+dx][y+dy] = dfs(x+dx, y+dy)
                        ans = max(ans, dp[x+dx][y+dy]+1)

            return ans

        for i, line in enumerate(matrix):
           for j, _ in enumerate(line):
              ans = max(ans, dfs(i, j))

        return ans

    def longestIncreasingPathfailed2(self, matrix):
        """
        :type matrix: List[List[int]]
        :rtype: int
        FAILED - TLE
        """
        if not matrix:
            return 0

        length = len(matrix)
        width = len(matrix[0])
        dir_vec = [(0, 1), (0, -1), (1, 0), (-1, 0)]

        lmat = []
        for i in range(length):
            lmat.append([1]*width)

        def flush():
            for x, line in enumerate(matrix):
                for y, _ in enumerate(line):
                    for dx, dy in dir_vec:
                        if x+dx<0 or x+dx>=length or y+dy<0 or y+dy>=width:
                           continue
                        elif matrix[x+dx][y+dy]>matrix[x][y]:
                            lmat[x+dx][y+dy] = max(lmat[x+dx][y+dy], 
                                                    lmat[x][y]+1)

        for i in range(max(length, width)):
            flush()

        ans = 0
        for i, line in enumerate(lmat):
            ans = max(ans, *line)

        return ans

    def longestIncreasingPathFailed(self, matrix):
        """
        :type matrix: List[List[int]]
        :rtype: int
        FAILED - TLE
        """
        if not matrix:
            return 0

        ans = 0
        length = len(matrix)
        width = len(matrix[0])
        dir_vec = [(0, 1), (0, -1), (1, 0), (-1, 0)]
        
        def dfs(x, y):
            ans = 1
            for dx, dy in dir_vec:
                # 判断越界
                if x+dx<0 or x+dx>=length or y+dy<0 or y+dy>=width:
                    continue
                elif matrix[x+dx][y+dy]>matrix[x][y]:
                    ans = max(ans, dfs(x+dx, y+dy)+1)

            return ans

        for i, line in enumerate(matrix):
           for j, _ in enumerate(line):
              ans = max(ans, dfs(i, j))

        return ans

nums = [
  [9,9,4],
  [6,6,8],
  [2,1,1]
]

nums = [
  [3,4,5],
  [3,2,6],
  [2,2,1]
]

obj = Solution()
print obj.longestIncreasingPath(nums)
