func rob(nums []int) int {
    prev2 := 0
    prev1 := 0

    for _, money := range nums {
        current := max(prev1, prev2+money)
        prev2 = prev1
        prev1 = current
    }

    return prev1
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}