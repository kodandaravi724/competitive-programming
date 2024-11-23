func dfs(i int, res *[]int, n int) {
    curDigit := 0
    for curDigit < 10 {
        if i*10 + curDigit <= n {
            *res = append(*res, i*10 + curDigit)
            dfs(i*10 + curDigit,res, n)
        } else {
            break
        }
        curDigit++
    }
}

func lexicalOrder(n int) []int {
    i := 1
    var res []int = []int{}
    for i < 10 {
        if i <= n {
            res = append(res, i)
            dfs(i, &res, n)
        } else {
            break
        }
        i++
    }
    return res
}