func max(i, j int) int {
    if i > j {
        return i
    }
    return j
}

func minimizedMaximum(n int, quantities []int) int {
    var maxQuantity int = -1
    for _, quantity := range(quantities) {
        maxQuantity = max(maxQuantity, quantity)
    }
    l := 1
    r := maxQuantity
    for l <= r {
        mid := l + (r-l)/2
        var minContainers float64 = 0
        for _, q := range(quantities) {
            minContainers += (math.Ceil(float64(q)/float64(mid)))
        }
        if minContainers > float64(n) {
            l = mid + 1
        } else {
            r = mid - 1
        }
    }
    return r+1
}