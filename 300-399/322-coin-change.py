from typing import List

class Solution:
    def coinChange(self, coins: List[int], amount: int) -> int:
        import math
        from collections import defaultdict

        if not amount:
            return 0

        array = defaultdict(lambda: defaultdict(lambda: math.inf))
        for i, coin in enumerate(coins):
            for j in range(1, amount + 1):
                mincoin = math.inf
                if j == coin:
                    mincoin = 1
                mincoin = min(mincoin, array[i - 1][j],
                              array[i][j - coin] + 1)
                array[i][j] = mincoin
        # for i, _ in enumerate(coins):
        #     for j in range(1, amount + 1):
        #         print(array[i][j], end=' ')
        #     print()
        result = array[i][j]
        return -1 if result is math.inf else result


a = Solution()
print(a.coinChange([1, 2, 5], 11))
# print(a.coinChange([2], 3))
