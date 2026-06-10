func rob(nums []int) int {
    rob1 := 0
    rob2 := 0

    for _, money := range nums {
        temp := max(rob1+money, rob2)
        rob1 = rob2
        rob2 = temp
    }

    return rob2
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}