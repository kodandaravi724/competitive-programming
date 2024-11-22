func setZeroes(matrix [][]int)  {
    var row []int = make([]int, len(matrix))
    var column []int = make([]int, len(matrix[0]))
    i := 0
    for i < len(matrix) {
        j := 0 
        for j < len(matrix[0]) {
            if matrix[i][j] == 0 {
                row[i] = 1
                column[j] = 1
            }
           j++ 
        }
        i++
    }
    i = 0
    for i < len(matrix) {
        j := 0
        for j < len(matrix[0]) {
            if row[i] == 1 || column[j] == 1 {
                matrix[i][j] = 0
            }
            j++
        }
        i++
    }
}