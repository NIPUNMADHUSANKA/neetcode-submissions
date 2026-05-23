func checkInclusion(s1 string, s2 string) bool {
	left := 0
	var s1Count [26]int
	var windowCount [26]int

	if len(s1) > len(s2){
		return false
	}

	for i := 0; i < len(s1); i++{
		s1Count[s1[i]-'a']++
		windowCount[s2[i]-'a']++
	}

	if s1Count == windowCount {
		return true
	}

	for right := len(s1); right < len(s2); right++{
		windowCount[s2[right]-'a']++
		windowCount[s2[left]-'a']--
		left++

		if s1Count == windowCount {
			return true
		}
	}
	return false
}
