func productExceptSelf(nums []int) []int {
	result := make([]int, len(nums))

	for i, _ := range nums{
		result[i] = 1
		for x, y := range nums{
			if i != x{
				result[i] = result[i] * y
			}
		}
	}

	return result
}
