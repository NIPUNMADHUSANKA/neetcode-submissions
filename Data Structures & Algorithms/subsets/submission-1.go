func subsets(nums []int) [][]int {
	var result [][]int
	var current []int

	var dfs func(index int)

	dfs = func(index int){
		if index >= len(nums){
			combination := append([]int{}, current...)
			result = append(result, combination)
			return
		}
		current = append(current,nums[index])
		dfs(index+1)
		current = current[:len(current)-1]
		dfs(index+1)
	}

	

	dfs(0)

	return result

}
