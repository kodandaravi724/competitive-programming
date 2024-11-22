type Pair struct {
    state int
    numRolls int
}
func min(i, j int) int {
    if i < j {
        return i
    }
    return j
}
func getCell (i int, n int) (int, int) {
    row := n-1-((i-1)/n)
    column := 0
    x := (i-1)/n + 1
    if x & 1 == 1 {
        column = ((i-1) % n)
    } else {
        column = (n-1) - ((i-1) % n)
    }
    return  row, column
}
func snakesAndLadders(board [][]int) int {
    n := len(board)
    v := make([]bool, n*n)
    q := []Pair{Pair{1,0,}}
    v[0] = true
    for len(q) > 0 {
        currentLevelSize := len(q)
        i := 0
        for i < currentLevelSize {
            currentState := q[0]
            if currentState.state == n*n {
                return currentState.numRolls
            }
            q = q[1:]
            roll := currentState.state + 1
            for roll <= min(n*n, currentState.state+6) {
                row, column := getCell(roll, n)
                if board[row][column] != -1 {
                    if !v[board[row][column]-1] {
                    q = append(q, Pair{board[row][column], currentState.numRolls+1})
                    v[board[row][column]-1] = true
                    }
                } else {
                    if !v[roll-1] {
                    q = append(q, Pair{roll, currentState.numRolls+1})
                    v[roll-1] = true
                    }
                }
                roll++
            }
            i++
        }
    }
    return -1
}