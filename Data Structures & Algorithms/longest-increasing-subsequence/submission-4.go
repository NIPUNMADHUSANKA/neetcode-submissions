func lengthOfLIS(nums []int) int {
    dp := make([]int, len(nums))

	for i := 0; i < len(nums); i++ {
		dp[i] = 1
	}

	for i:= 0; i < len(nums); i++{
		for j:= 0; j < i; j++{
			if nums[j] < nums[i]{
				dp[i] = max(dp[i],dp[j]+1)
			}
		}
	}
	return maxList(dp)
}


func max(a, b int)int{
	if a > b{
		return a
	}else{
		return b
	}
}

func maxList(dp []int)int{
	max := 0
	for _, i := range dp{
		if i > max{
			max = i
		}
	}
	return max
}

