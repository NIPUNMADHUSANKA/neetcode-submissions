func singleNumber(nums []int) int {

	result := 0

	for _, i := range nums{
		result = result ^ i
	}
	return result
}
