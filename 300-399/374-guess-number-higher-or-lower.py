# The guess API is already defined for you.
# @param num, your guess
# @return -1 if my number is lower, 1 if my number is higher, otherwise return 0
# def guess(num):

# 二分查找

result = 1
def guess(num):
    if result > num:
        return 1
    elif num == result:
        return 0
    else:
        return -1


class Solution(object):
    def guessNumber(self, n):
        """
        :type n: int
        :rtype: int
        """
        low, high = 1, n
        while True:
            num = (low + high) // 2
            result = guess(num)
            if result == 0:
                return num
            elif result == 1:
                low = num + 1
            else:
                high = num


s = Solution()
print(s.guessNumber(1))
