func lengthOfLongestSubstring(s string) int {
	left := 0
	set  := make(map[byte]int)
	result := 0

	for right := 0; right < len(s); right++{
		set[s[right]]++

		for set[s[right]] > 1 {
			set[s[left]]--
			left++
		}

		result = max(result, right - left + 1)
	} 
	return result
}

func max(x, y int)int{
	if x > y{
		return x
	}
	return y
}