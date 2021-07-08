from typing import List


class Solution:
    def spiralOrder(self, matrix: List[List[int]]) -> List[int]:
        if not matrix:
            return []
        return list(matrix[0]) + self.spiralOrder(list(zip(*matrix[1:]))[::-1])


if __name__ == '__main__':
    s = Solution()
    print(s.spiralOrder([
     [ 1, 2, 3 ],
     [ 4, 5, 6 ],
     [ 7, 8, 9 ]
    ]))
