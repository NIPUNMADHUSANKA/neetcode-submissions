func maxArea(heights []int) int {
    left := 0
    right := len(heights) - 1
    result := 0

    for left < right {
        result = max(result, min(heights[left], heights[right])*(right-left))

        if heights[left] <= heights[right]{
            left++
        }else{
            right--
        }
    }
    return result

}


func max(a,b int)int{
    if a >= b {
        return a
    }else{
        return b
    }
}

func min(a,b int)int{
    if a >= b {
        return b
    }else{
        return a
    }
}