package main

// https://leetcode.com/problems/n-repeated-element-in-size-2n-array/discuss/298319/0ms-java-solution-without-map-set-or-array

func repeatedNTimes(A []int) int {
    for i := 2; i < len(A); i++ {
        if (A[i] == A[i-1]) || (A[i] == A[i-2]) {
            return A[i]
        }
    }
    return A[0]
}