func characterReplacement(s string, k int) int {
	left := 0
	maxFre := 0
	result := 0
	counter := make(map[byte]int)

	for right := 0; right < len(s); right++{
		counter[s[right]]++
		
		if counter[s[right]] > maxFre{
			maxFre = counter[s[right]]
		}

		for right-left+1-maxFre > k{
			counter[s[left]]--
			left++
		}

		result = max(result, right-left+1)
	}
	return result
}


func max(a, b int)int{
	if a > b {
		return a
	} else{
		return b
	}
}