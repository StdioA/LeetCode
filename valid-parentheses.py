# coding: utf-8

class Solution:
    # @return a boolean
    def isValid(self, s):
        brackNum = {
            '(' : 1,
            ')' : -1,
            '{' : 2,
            '}' : -2,
            '[' : 3,
            ']' : -3
            }
        nums = [brackNum[x] for x in s]
        stack = []
        for b in nums:
            if b > 0:
                stack.append(b)
            else:
                if stack == []:
                    return False
                if b+stack[-1] == 0:
                    stack.pop()
                else:
                    return False
        if stack != []:
            return False
        return True