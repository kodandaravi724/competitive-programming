func maximumCoins(coins [][]int, k int) int64 {
    i := 0
    j := 0
    res := 0
    sort.Slice(coins, func (i, j int) bool {
        return coins[i][0] < coins[j][0]
    })
    n := len(coins)
    current := 0
    res = 0
    for i < n && j < n {
        // fmt.Println(coins)
        if coins[j][0] + k > coins[i][1] {
            current += (coins[i][1]-coins[i][0]+1)*coins[i][2]
            res = max(res, current)
            i++
        } else {
            partial := max(0, coins[j][0]+k-coins[i][0])*coins[i][2]
            res = max(res, current+partial)
            current -= (coins[j][1]-coins[j][0]+1)*coins[j][2]
            j++
        }
    }
    i = n - 1
    j = n - 1
    current = 0
    for i >= 0 && j >= 0 {
        if coins[j][1] - k < coins[i][0] {
            current += (coins[i][1]-coins[i][0]+1)*coins[i][2]
            res = max(res, current)
            i--
        } else {
            partial := max(0, (-coins[j][1]+k+coins[i][1])*coins[i][2])
            // fmt.Println(partial, i, j, coins, -coins[j][1]+k+coins[i][1])
            res = max(res, current+partial)
            current -= (coins[j][1]-coins[j][0]+1)*coins[j][2]
            j--
        }
    }
    return int64(res)
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}