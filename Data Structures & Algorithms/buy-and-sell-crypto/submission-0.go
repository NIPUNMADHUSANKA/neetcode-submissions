func maxProfit(prices []int) int {
	left := 0
	maxProfit := 0

	for right := 0; right < len(prices); right++{
		maxProfit = max(maxProfit, prices[right]-prices[left])

		for prices[right] < prices[left]{
			left++
		}
	}
	return maxProfit
}


func max(a, b int)(int){
	if a > b {
		return a
	}
	return b
}
