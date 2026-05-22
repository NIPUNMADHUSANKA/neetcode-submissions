func lengthOfLongestSubstring(s string) int {
	left := 0
	result := 0
	seen := make(map[byte]int)

	for right := 0; right < len(s); right++{
		if idx, ok := seen[s[right]]; ok {
			left = max(left, idx+1)
		}
		seen[s[right]] = right
		result = max(result, right-left+1)
	}
	return result
}

func max(a,b int)int{
	if a > b{
		return a
	}else{
		return b
	}
}