func longestConsecutive(nums []int) int {
    set := make(map[int]struct{})

    for _, n := range nums {
        set[n] = struct{}{}
    }

    longest := 0

    for n := range set {
        // Only start counting if n is the beginning of a sequence
        if _, exists := set[n-1]; !exists {
            current := n
            length := 1

            for {
                if _, exists := set[current+1]; exists {
                    current++
                    length++
                } else {
                    break
                }
            }

            if length > longest {
                longest = length
            }
        }
    }

    return longest
}