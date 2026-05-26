func findMin(nums []int) int {
	if len(nums) == 1{
		return nums[0]
	}else{
		leftArr := nums[:len(nums)/2]
		rightArr := nums[len(nums)/2:]

		left := findMin(leftArr)
		right := findMin(rightArr)

		if left < right {
			return left
		}else{
			return right
		}
	}
}

