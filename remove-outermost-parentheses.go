package main

import "strings"

func removeOuterParentheses(S string) string {
    levels := make([]int, len(S))
    level := 0
    for i, p := range S {
        if p == '(' {
            level++
            levels[i] = level
        } else {
            levels[i] = level
            level--
        }
    }
    var b strings.Builder
    for i, r := range levels {
        if r > 1 {
            b.WriteByte(S[i])
        }
    }
    return b.String()
}