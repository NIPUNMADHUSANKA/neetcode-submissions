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


/*
func maxSubArray(nums []int) (int, int, int) {
	currentMax := nums[0]
	globalMax := nums[0]

	start := 0
	end := 0
	tempStart := 0

	for i := 1; i < len(nums); i++ {

		// Start a new subarray from i
		if nums[i] > currentMax+nums[i] {
			currentMax = nums[i]
			tempStart = i
		} else {
			// Continue existing subarray
			currentMax += nums[i]
		}

		// Found a better maximum
		if currentMax > globalMax {
			globalMax = currentMax
			start = tempStart
			end = i
		}
	}

	return globalMax, start, end
} */