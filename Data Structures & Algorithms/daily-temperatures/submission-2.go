func dailyTemperatures(temperatures []int) []int {
	result := make([]int, len(temperatures))
	stack := []int{}

	for i, temp := range temperatures{
		for len(stack) > 0 && temp > temperatures[stack[len(stack)-1]]{
			prev := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			result[prev] = i - prev	
		}
		stack = append(stack, i)
	}
	return result
}
