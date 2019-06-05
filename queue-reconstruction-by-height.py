from typing import List


class Solution:
    def reconstructQueue(self, people: List[List[int]]) -> List[List[int]]:
        people.sort()
        result = [None] * len(people)
        for h, k in people:
            c, i = 0, 0
            while i < len(result):
                pos = result[i]
                # 找空位
                if c >= k and pos is None:
                    break
                # 给比 h 高的人留出位置
                if pos is None or pos[0] >= h:
                    c += 1
                i += 1
            result[i] = [h, k]
        return result

s = Solution()
print(s.reconstructQueue([[7, 0], [4, 4], [7, 1], [5, 0], [6, 1], [5, 2]]))
