func twoSum(nums []int, target int) []int {
    m := make(map[int]int, len(nums))
    for i,n := range nums {
        //check
        if v, ok := m[n]; ok {
            return []int{v, i}
        } else {
            //add
            m[target - n] = i
        }
    }
    return []int{}
}