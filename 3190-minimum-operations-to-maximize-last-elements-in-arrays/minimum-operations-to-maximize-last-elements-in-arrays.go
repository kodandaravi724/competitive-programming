func numOperations(nums1 []int, nums2 []int) int {
    i := 0
    n := len(nums1)
    swaps := 0
    for i < n {
        if nums1[i] <= nums1[n-1] && nums2[i] <= nums2[n-1] {
            i++
        } else if(nums1[i] <= nums2[n-1] && nums2[i] <= nums1[n-1]) {
            swaps++
            i++
        } else {
            return -1
        }
    }
    return swaps
}

func min(a int, b int) int {
    if a < b {
        return a
    }
    return b
}

func minOperations(nums1 []int, nums2 []int) int {
    n := len(nums1)
    c1 := numOperations(nums1, nums2)
    nums1[n-1], nums2[n-1] = nums2[n-1], nums1[n-1]
    c2 := 1 + numOperations(nums1, nums2)
    return min(c1, c2)
}