func findMin(nums []int) int {
	left, right := 0, len(nums)-1

	for left < right{
		mid := left + (right-left)/2
		if nums[mid] > nums[right]{
			left = mid + 1
		}else{
			right = mid
		}
	}
	return nums[left] 
}

/*
A rotated sorted array must come from a fully sorted array with only one breaking point.
*/