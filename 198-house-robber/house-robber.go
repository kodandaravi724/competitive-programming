func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func rob(nums []int) int {
    var dp []int = make([]int, len(nums)+1)
    n := len(nums)
    dp[n-1] = nums[n-1]
    i := n-2
    for i >= 0 {
        dp[i] = max(nums[i]+dp[i+2], dp[i+1])
        i--
    }
    return dp[0]
}