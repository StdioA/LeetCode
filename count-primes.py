# coding: utf-8

# UNFINISHED!!!

class Solution(object):
    def countPrimes(self, n):
        """
            :rtype: int
            :type n: int
        """

        if n < 2:
            return 0
        primes = [2]
        for num in range(3, n):
            for p in primes:
                if num%p == 0:
                    break
            else:
                primes.append(num)
        return len(primes)

a = Solution()
print(a.countPrimes(50))
