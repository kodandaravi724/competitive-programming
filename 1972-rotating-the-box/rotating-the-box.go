func move(row []byte) {
    i := len(row) - 1
    numRocks := 0
    numEmpty := 0
    rightBound := len(row)
    for i >= 0 {
        for i>=0 && row[i] != '*' {
            if row[i] == '.' {
                numEmpty++
            } else {
                numRocks++
            }
            i--
        }
        if numEmpty > 0 && numRocks > 0 {
            k := i+1
            for k < rightBound {
                if numEmpty > 0 {
                    row[k] = '.'
                    numEmpty--
                } else {
                    row[k] = '#'
                    numRocks--
                }
                k++
            }
        } else {
            numEmpty = 0
            numRocks = 0
        }
        rightBound = i
        i--
    }
}

func rotateTheBox(box [][]byte) [][]byte {
    i := 0
    var result [][]byte = make([][]byte, len(box[0]))
    for i < len(box) {
        move(box[i])
        i++
    }
    i = 0
    for i < len(box[0]) {
        result[i] = make([]byte, len(box))
        i++
    }
    i = 0
    m := len(box)-1
    // n := len(box[0])-1
    for i < len(box[0]) {
        j := 0
        for j < len(box) {
            result[i][j] = box[m-j][i]
            j++
        }
        i++
    }
    return result
}