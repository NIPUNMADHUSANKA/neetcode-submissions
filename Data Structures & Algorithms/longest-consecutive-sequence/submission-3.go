func longestConsecutive(nums []int) int {
	set := make(map[int]struct{})

	for _, v := range nums{
		set[v] = struct{}{}
	}

	longest := 0

	for i := range set{
		if _, exist := set[i-1]; !exist{
			lenght := 1
			for true{
				if _, exist := set[i+lenght]; exist{
					lenght++
				}else{
					break
				}
			}

			if lenght > longest{
				longest = lenght
			}
		}
	}

	return longest
}
