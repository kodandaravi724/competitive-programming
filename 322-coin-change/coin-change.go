func coinChange(coins []int, amount int) int {
    var dp [][]int = make([][]int, len(coins)+1)
    i := 0
    n := len(coins)
    for i <= n {
        dp[i] = make([]int, amount + 1)
        i++
    }
    i = 1
    for i <= amount {
        dp[0][i] = -1
        i++
    }
    i = 1
    for i <= n {
        j := 1
        for j <= amount {
            if coins[i-1] <= j {
                if dp[i-1][j] == -1 {
                    if dp[i][j-coins[i-1]] == -1 {
                        dp[i][j] = -1
                    } else {
                        dp[i][j] = 1+dp[i][j-coins[i-1]]
                    }
                } else {
                    if dp[i][j-coins[i-1]] == -1 {
                        dp[i][j] = dp[i-1][j]
                    } else {
                    dp[i][j] = min(1+dp[i][j-coins[i-1]], dp[i-1][j])
                    }
                }
            } else {
                dp[i][j] = dp[i-1][j]
            }
            j++
        }
        i++
    }
    // for ind, d := range(dp) {
    //     fmt.Println(ind, d)
    // }
    return dp[len(coins)][amount]
}