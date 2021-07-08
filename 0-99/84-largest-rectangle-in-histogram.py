# coding: utf-8

class Solution(object):
    def largestRectangleArea(self, height):
        """
        :type height: List[int]
        :rtype: int
        """
        height.append(0)
        max_area = 0
        stack = []
        i = 0

        while i < len(height):
            if (not stack) or (height[stack[-1]] < height[i]):
                stack.append(i)
                i += 1
            else:
                t = stack.pop()
                if not stack:
                    max_area = max(max_area, height[t] * (i));
                else:
                    max_area = max(max_area, height[t]*(i - stack[-1] - 1))

            # print [(q, height[q]) for q in stack], max_area

        return max_area

a = Solution()
print a.largestRectangleArea([1, 2, 2])

