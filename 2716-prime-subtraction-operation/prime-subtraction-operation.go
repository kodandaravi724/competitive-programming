func primeSubOperation(nums []int) bool {
    isPrime := func(n int) bool {
        if n <= 1 {
            return false
        }
        for i := 2; i*i <= n; i++ {
            if n%i == 0 {
                return false
            }
        }
        return true
    }

    primes := []int{}
    for i := 2; i < 1000; i++ {
        if isPrime(i) {
            primes = append(primes, i)
        }
    }

    // Iterate over the array from the second last to the first element
    for i := len(nums) - 2; i >= 0; i-- {
        if nums[i] >= nums[i+1] {
            canSubtract := false
            for _, p := range primes {
                if p >= nums[i] {
                    break
                }
                if nums[i] - p < nums[i+1] {
                    nums[i] -= p
                    canSubtract = true
                    break
                }
            }
            if !canSubtract {
                return false
            }
        }
    }
    return true
}