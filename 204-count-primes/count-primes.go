func countPrimes(n int) int {
    var prime []bool = []bool{}
    p := 2
    for i := 0; i <= n; i++ {
        prime = append(prime, true)
    }
    for p*p <= n {
        if prime[p] {
            k := p*p
            for k <= n {
                prime[k] = false
                k += p
            }
        } 
        p += 1
    }
    count := 0
    for i := 2; i < n; i++ {
        if prime[i] {
            count++
        }
    }
    return count
}