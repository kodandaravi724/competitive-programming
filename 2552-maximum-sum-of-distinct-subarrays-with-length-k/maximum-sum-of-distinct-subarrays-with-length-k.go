func max(i, j int) int {
    if i > j {
        return i
    }
    return j
}

func maximumSubarraySum(nums []int, k int) int64 {
    var numMap map[int]bool = make(map[int]bool)
    var windowStart, windowEnd, maxSum, runningSum int = 0, 0, 0, 0

    for windowEnd < len(nums) {
        runningSum += nums[windowEnd]
        _, ok := numMap[nums[windowEnd]]
        if ok {
            for nums[windowStart] != nums[windowEnd] {
                delete(numMap, nums[windowStart])
                runningSum -= nums[windowStart]
                windowStart++
            }
            delete(numMap, nums[windowStart])
            runningSum -= nums[windowStart]
            windowStart++
            numMap[nums[windowEnd]] = true
        } else {
            numMap[nums[windowEnd]] = true
        }
        if (windowEnd - windowStart +1) == k {
            maxSum = max(maxSum, runningSum)
            runningSum -= nums[windowStart]
            delete(numMap, nums[windowStart])
            windowStart++
        }
        windowEnd++
    }

    return int64(maxSum)
}