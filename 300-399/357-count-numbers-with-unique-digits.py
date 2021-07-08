class Solution:
    def countNumbersWithUniqueDigits(self, n: int) -> int:
        if n == 0:
            return 1
        # 1(L0) + (10 - 1)(L1, without 0) + 9 * 9(L2, including 0)
        # + 9 * C(9, 2)(L3) + 9 * C(9, 3)(L4)
        numbers = 9
        x = 9
        for i in range(1, min(n, 11)):
            numbers += 9 * x
            x *= 9 - i
        return numbers + 1


s = Solution()
print(s.countNumbersWithUniqueDigits(11))
