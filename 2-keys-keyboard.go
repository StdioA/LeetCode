func minSteps(n int) int {
    ans, d := 0, 2
    for n > 1 {
        for n % d == 0 {
            ans += d
            n /= d
        }
        d += 1
    }
    return ans
}
