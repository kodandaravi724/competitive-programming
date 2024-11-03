func trailingZeroes(n int) int {
    if n < 0 {
        return -1
    }
    var trailingZeros = 0
    for n >= 5 {
        n = n / 5
        trailingZeros += n
    }
    return trailingZeros
}