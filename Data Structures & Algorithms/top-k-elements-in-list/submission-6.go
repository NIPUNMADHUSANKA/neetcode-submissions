
func topKFrequent(nums []int, k int) []int {
    freq := make(map[int]int)
    bucket := make([][]int, len(nums)+1)
    result := []int{}


    for _, val := range nums{
        freq[val]++
    }

    for num, count := range freq {
        bucket[count] = append(bucket[count], num)
    }

    for i:= len(bucket)-1; i >= 0 && len(result) < k; i--{
        result = append(result, bucket[i]...)
    }
    return result[:k]

}
                                                                                                        