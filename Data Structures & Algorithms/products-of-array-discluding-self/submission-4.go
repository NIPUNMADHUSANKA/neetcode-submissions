func productExceptSelf(nums []int) []int {
  result := make([]int, len(nums))

  prefix := 1
  postfix := 1


  for i, j := range nums{
	result[i] = prefix
	prefix = prefix * j
  }

  for i := len(result)-1; i>=0; i--{
	result[i] = result[i] * postfix
	postfix = postfix * nums[i]
  }
  return result

}
