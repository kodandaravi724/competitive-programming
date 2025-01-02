func vowelStrings(words []string, queries [][]int) []int {
    num := make([]int, len(words))
    i := 0
    for i < len(words) {
        toAdd := isvowel(words[i][0]) && isvowel(words[i][len(words[i])-1])
        if i == 0 {
            if toAdd {
                num[i] = 1
            }
        } else {
            if toAdd {
                num[i] = 1 + num[i-1]
            } else {
                num[i] = num[i-1]
            }
        }
        i++
    }
    res := make([]int, len(queries))
    for ind, q := range queries {
        if q[0] == 0 {
            // fmt.Println(q)
            res[ind] = num[q[1]]
        } else {
            res[ind] = num[q[1]] - num[q[0]-1]
        }
    }
    return res
}

func isvowel(a byte) bool {
    return a == 'a' || a == 'e' || a == 'i' || a == 'o' || a == 'u'
}