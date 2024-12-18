func isPossible(ribbons []int, k int, l int) bool {
    var s int = 0
    for _, ribbon := range(ribbons) {
        s += ribbon/l
    }
    return k <= s
}

func sum(ribbons []int) int {
    s := 0
    i := 0
    for i < len(ribbons) {
        s += ribbons[i]
        i+=1
    }
    return s
}

func maxLength(ribbons []int, k int) int {
    r := -1
    for _, ribbon := range(ribbons) {
        if r < ribbon {
            r = ribbon
        }
    }

    l := 1
    if k > sum(ribbons) {
        return 0
    }

    for l < r {
        mid := (l+r+1)/2
        if isPossible(ribbons, k, mid) {
            l = mid
        } else {
            r = mid - 1
        }
    }
    return l
}