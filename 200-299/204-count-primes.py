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


class Solution(object):
    def countPrimes(self, n):
        """
            :rtype: int
            :type n: int
        """
        if n<3:
            return 0

        import math

        map_ = [[m, 1] for m in range(3, n, 2)]
        count = 1

        for index, obj in enumerate(map_):
            num, prime = obj
            if prime:
                count += 1
                print("prime", num)

                if num > math.sqrt(n)+1:
                    continue

                for i in range(1, (n//num+1)//2):
                    if index+num*i > n//2:
                        break
                    print("filter", map_[index+num*i][0])


        # print([m for m, p in map_ if p])
        return count


a = Solution()
print(a.countPrimes(1500000))
