func sum(nums []int) int {
    sum := 0
    for _, num := range(nums) {
        sum += num
    }
    return sum
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func getLargestOutlier(nums []int) int {
    totalSum := sum(nums)

    numMap := make(map[int]int)

    for index, num := range(nums) {
        numMap[num] = index
    }

    i := 0
    n := len(nums)
    largestOutlier := -10001
    for i < n {
        remainingSum := totalSum - nums[i]
        if remainingSum % 2 == 0  {
            sumV := remainingSum / 2
            index, ok := numMap[sumV]
            if ok {
                if index != i {
                    largestOutlier = max(largestOutlier, nums[i])
                }
            }
        }
        i++
    }
    return largestOutlier
}