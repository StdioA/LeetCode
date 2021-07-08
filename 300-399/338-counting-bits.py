# coding: utf-8

class Solution(object):
    def countBits(self, num):
        """
        :type num: int
        :rtype: List[int]
        """
        bits = []
        for i in range(num+1):
            bits.append(len(filter(lambda x: x=="1", bin(i)[2:])))

        return bits

    def countBits2(self, num):
        """
        递推方式求解
        0 1 12 1223 12232334
        前一半跟上一组一样，后一半等于上一组+1
        """
        bitlen = len(bin(num)) - 1

        bits = [0] * (2**bitlen)
        bits[1] = 1

        for i in range(2, bitlen+1):
            bits[2**(i-1):2**(i-1)+2**(i-2)] = bits[2**(i-2):2**(i-1)]
            bits[2**(i-1)+2**(i-2):2**(i)] = map(lambda x: x+1, bits[2**(i-2):2**(i-1)])

        return bits[:num+1]

a = Solution()
print a.countBits2(20)
