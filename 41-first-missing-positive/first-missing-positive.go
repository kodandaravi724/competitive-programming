func firstMissingPositive(nums []int) int {
    i := 0
    n := len(nums)
    for i < len(nums) {
        if nums[i]-1 == i {
            i++
            continue
        } else {
            for nums[i] >= 1 && nums[i] <= n && (nums[i] -1 != i){
                k := nums[i]
                if nums[k-1] == nums[i] {
                    break
                }
                nums[i], nums[k-1] = nums[k-1], nums[i]
            }
            i++
        }
    }
    i = 0
    for i < n {
        if nums[i]-1 != i {
            return i+1
        }
        i++
    }
    return i+1
}