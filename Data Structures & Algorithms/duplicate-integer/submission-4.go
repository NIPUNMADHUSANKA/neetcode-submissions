func hasDuplicate(nums []int) bool {
    hashList := make(map[int]struct{})

    for _, val := range nums{
        if _, exist := hashList[val]; exist{
            return true
        }
        hashList[val] = struct{}{}
    }
    return false
}
