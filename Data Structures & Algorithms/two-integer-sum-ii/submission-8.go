func twoSum(numbers []int, target int) []int {
    left := 0
    right := len(numbers) - 1

    for left < right{
        num := numbers[left] + numbers[right]
        if  num > target {
            right--
        }else if num == target{
            return []int{left+1, right+1}
        }else{
            left++
        }
    }
    return []int{}
}
