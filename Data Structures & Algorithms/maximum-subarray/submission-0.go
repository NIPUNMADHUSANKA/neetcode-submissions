func maxSubArray(nums []int) int {
    currentMax := nums[0]
    globalMax := nums[0]

    for i := 1; i < len(nums); i++ {
        currentMax = myMax(nums[i], currentMax+nums[i])
        globalMax = myMax(globalMax, currentMax)
    }

    return globalMax
}


func myMax(a,b int)int{
    if a > b{
        return a
    }else{
        return b
    }
}