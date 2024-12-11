func maximumBeauty(nums []int, k int) int {
            max := 300002

        var sweep []int = make([]int, max)

        for _, num := range(nums) {
            sweep[num-k+100000] += 1
            sweep[num+k+100001] -= 1
        }
        ans := -1
        t := 0
        for _, e := range(sweep) {
            t += e
            if ans < t {
                ans = t
            }
        }
        return ans
}