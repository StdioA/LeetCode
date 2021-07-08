# coding: utf-8


class Solution(object):
    def maximalSquare(self, matrix):
        """
        :type matrix: List[List[str]]
        :rtype: int
        """
        if not matrix:
            return 0

        m, n = len(matrix), len(matrix[0])
        st = []
        for i in range(m):
            k = []
            for j in range(n):
                k.append([0, 0])
            st.append(k)                                                        # 每个点所能构成长方形的最大宽度和高度

        max_ = 0
        for i in range(m):
            for j in range(n):
                if matrix[i][j]=='1':
                    if i==0:
                        st[i][j][0] = 1
                    else:
                        st[i][j][0] = min(st[i-1][j][0], st[i-1][j-1][0])+1

                    if j==0:
                        st[i][j][1] = 1
                    else:
                        st[i][j][1] = min(st[i][j-1][1], st[i-1][j-1][1])+1

                    if st[i][j][0] == st[i][j][1]:                              # 判断是否为正方形
                        max_ = max(max_, st[i][j][0]**2)

        for i in range(m):
            print([[t[0], t[1], matrix[i][index]] for index, t in enumerate(st[i])])
        return max_

obj = Solution()
print(obj.maximalSquare(["101001110","111000001","001100011","011001001","110110010","011111101","101110010","111010001","011110010","100111000"]))
# print(obj.maximalSquare(['0']))
