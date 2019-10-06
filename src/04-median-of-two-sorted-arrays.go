/*
O(n+m) solution. Should be O(log(n+m)). Needs improvement https://www.geeksforgeeks.org/median-two-sorted-arrays-different-sizes-ologminn-m/
*/

func getMedian(nums []int) float64 {
    size := len(nums)
    if size % 2 == 0 {
        return float64(nums[size/2-1] + nums[size/2]) / 2
    }
    return float64(nums[size/2])
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    ptr1, ptr2, index := 0, 0, 0 
    size := len(nums1) + len(nums2)
    merged := make([]int, size)
    for ptr1 < len(nums1) && ptr2 < len(nums2) {
        if nums1[ptr1] < nums2[ptr2] {
            merged[index] = nums1[ptr1]
            ptr1 = ptr1 + 1
        } else {
            merged[index] = nums2[ptr2]
            ptr2 = ptr2 + 1
        }
        index = index + 1
    }
    
    for ptr1 < len(nums1) {
        merged[index] = nums1[ptr1]
        ptr1 = ptr1 + 1
        index = index + 1
    }
    
    for ptr2 < len(nums2) {
        merged[index] = nums2[ptr2]
        ptr2 = ptr2 + 1
        index = index + 1
    }

    return getMedian(merged)
}
