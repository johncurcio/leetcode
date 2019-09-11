func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func max(a, b int) int {
    if a < b {
        return b
    }
    return a
}

// time O(n), space O(1)
func maxArea(height []int) int {
    maxarea := 0
    beg, end := 0, len(height)-1
    for beg < end {
        width := end-beg
        maxarea = max(maxarea, min(height[end], height[beg]) * width)
        if height[beg] < height[end] {
            beg++
        } else {
            end--
        }
    }
    return maxarea
}