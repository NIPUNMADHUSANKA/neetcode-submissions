func combinationSum(nums []int, target int) [][]int {
    var result [][]int
	var current []int

	var dfs func(index int, remaining int)

	dfs = func(index int, remaining int){
		if remaining == 0 {
			combination := append([]int{}, current...)
			result = append(result, combination)
			return
		}
		if index >= len(nums) || remaining < 0{
			return
		}

		current = append(current,nums[index])
		dfs(index, remaining - nums[index])

		current = current[:len(current)-1]
		
		dfs(index+1, remaining)

	}

	dfs(0, target)

	return result
}
