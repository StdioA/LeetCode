# coding: utf-8

#del, pop(), or remove() are not allowed.
# 评测机让人感觉好奇怪，提交了27次都没过，找的别人的答案
class Solution:
    # @param a list of integers
    # @return an integer
    def removeDuplicates(self, A):
        if not A:
            return 0
        end = len(A)
        read = 1
        write = 1
        while read < end:
            if A[read] != A[read-1]:
                A[write] = A[read]
                write += 1
            read += 1
        return write

a = Solution()
print a.removeDuplicates([1,2,3,3,4,5,5])
