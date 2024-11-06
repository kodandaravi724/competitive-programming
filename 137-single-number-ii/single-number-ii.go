func singleNumber(nums []int) int {
    var a, b int = 0, 0
    for _, element := range(nums) {
        a = (a^element) &(^b)
        b = (b^element) &(^a)
    }
    return a
 }