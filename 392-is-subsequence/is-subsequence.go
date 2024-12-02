func isSubsequence(s string, t string) bool {
    i := 0
    j := 0
    m := len(s)
    n := len(t)
    for i < m && j < n {
        if s[i] == t[j] {
            i++
            j++
        } else {
            j++
        }
    }
    return i == m
}