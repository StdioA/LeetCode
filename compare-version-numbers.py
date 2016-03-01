# coding: utf-8

class Solution:
    # @param version1, a string
    # @param version2, a string
    # @return an integer
    def compareVersion(self, version1, version2):
        if version1 == version2:
            return 0
        v1 = [int(x) for x in version1.split('.')]
        v2 = [int(x) for x in version2.split('.')]
        l1 = len(v1)
        l2 = len(v2)
        lmin = min(l1, l2)
        for x in range(lmin):
            if v1[x] > v2[x]:
                return 1
            elif v1[x] < v2[x]:
                return -1
        if l1 > l2:
            for x in range(l1-l2):
                if v1[l2+x] != 0:
                    return 1
        elif l1 < l2:
            for x in range(l2-l1):
                if v2[l1+x] != 0:
                    return -1
        return 0