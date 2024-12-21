func rob(nums []int) int {
    if len(nums) == 1 {
        return nums[0]
    }
    return max(rob_noncyclic(nums[1:]), rob_noncyclic(nums[:len(nums)-1]))
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func rob_noncyclic(nums []int) int {
    fmt.Println(nums)
    var dp []int = make([]int, len(nums)+1)
    n := len(nums)
    if n == 0 {
        return 0
    }
    dp[n-1] = nums[n-1]
    i := n-2
    for i >= 0 {
        dp[i] = max(nums[i]+dp[i+2], dp[i+1])
        i--
    }
    return dp[0]
}