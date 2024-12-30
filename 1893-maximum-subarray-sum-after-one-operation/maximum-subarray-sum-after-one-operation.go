func getMax(arr []int) int {
    var max int = arr[0]
    i := 1
    n := len(arr)
    for i < n {
        if max < arr[i] {
            max = arr[i]
        }
        i++
    }
    return max
}

func getMin(arr []int) int {
    var min int = arr[0]
    i := 1
    n := len(arr)
    for i < n {
        if min > arr[i] {
            min = arr[i]
        }
        i++
    }
    return min
}

func getSum(arr []int, n int) int {
    var sum int = arr[0]
    i := 1
    for i < n {
        sum += arr[i]
        i++
    }
    return sum
}


func maxSumAfterOperation(nums []int) int {
    // var max, min, sum int = getMax(nums), getMin(nums), getSum(nums, len(nums))
    // if max*max - max > min*min-2*min {
    //     return sum + max*max - max 
    // }
    // // fmt.Println(sum, min)
    // return sum + min*min-min
    var maxLeft, maxRight []int = make([]int, len(nums)), make([]int, len(nums))
    var n int = len(nums)

    i := 1
    for i < n {
        maxLeft[i] = max(0, maxLeft[i-1]+nums[i-1])
        i++
    }

    i = n - 2
    for i >= 0 {
        maxRight[i] = max(0, maxRight[i+1]+nums[i+1])
        i--
    }

    ans := 0
    i = 0
    for i < n {
        ans = max(nums[i]*nums[i]+maxRight[i]+maxLeft[i], ans)
        i++
    }
    
    return ans
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}