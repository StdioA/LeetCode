# coding: utf-8

class Solution:
    # @param A  a list of integers
    # @param m  an integer, length of A
    # @param B  a list of integers
    # @param n  an integer, length of B
    # @return nothing
    def merge(self, A, m, B, n):
        totalLength = m+n
        while m > 0 and n >0:
            if A[m-1] >= B[n-1]:
                A[totalLength - 1] = A[m - 1]
                m -= 1
                totalLength -= 1
            elif B[n - 1] >= A[m -1]:
                A[totalLength - 1] = B[n-1]
                n -= 1
                totalLength -= 1
        if n > 0:
            A[0:n] = B[0:n]

# testcode太奇怪，实在A不掉了…