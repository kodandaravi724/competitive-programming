func getCombinations(candidates []int, i int, target int, res *[][]int, cur []int) {
    if target == 0 {
        *res = append(*res, append([]int{}, cur...))
        return
    }
    if target < 0 {
        return
    }
    for i < len(candidates) {
        cur = append(cur, candidates[i])
        getCombinations(candidates, i, target - candidates[i], res, cur)
        cur = cur[:len(cur)-1]
        i++
    }
}

func combinationSum(candidates []int, target int) [][]int {
    var combinations [][]int = [][]int{}
    getCombinations(candidates, 0, target, &combinations, []int{})
    return combinations
}