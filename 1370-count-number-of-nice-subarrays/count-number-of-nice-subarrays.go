func numberOfSubarrays(nums []int, k int) int {
    var l, m, r, n, oddCount, niceArrays int = 0, 0, 0, len(nums), 0, 0

    for r < n {
        if nums[r]&1 == 1 {
            oddCount++
        }
        for oddCount > k {
            if nums[l]&1 == 1 {
                oddCount--
            }
            l++
            m = l
        }
        if oddCount == k {
            for nums[m]&1 == 0 {
                m++
            }
            niceArrays += (m-l+1)
        }
        r++
    }

    return niceArrays
}