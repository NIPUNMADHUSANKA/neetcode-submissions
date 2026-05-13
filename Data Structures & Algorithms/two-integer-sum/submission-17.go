func twoSum(nums []int, target int) []int {
    m := make(map[int]int)

    for i, v := range(nums){
        check := target - v
        if j, ok := m[check]; ok{
            return []int{j,i}
        }
        m[v] = i
    }
    return []int{}
}
