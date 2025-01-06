func maximumWhiteTiles(tiles [][]int, carpetLen int) int {
    var i, j int = 0, 0
    sort.Slice(tiles, func(i, j int) bool {
        return tiles[i][0] < tiles[j][0]
    })
    res := 0
    cover := 0
    for i < len(tiles) &&  res < carpetLen {
        // fmt.Println(i,j, res, cover)
        if  tiles[j][0] + carpetLen > tiles[i][1] {
            cover += (tiles[i][1] - tiles[i][0] + 1)
            // fmt.Println((tiles[i][1] - tiles[i][0] + 1), i)
            res = max(res, cover)
            i++
        } else {
            partial := max(0, tiles[j][0]+carpetLen-tiles[i][0])
            // cover += partial
            res = max(res, cover+partial)
            cover -= (tiles[j][1] - tiles[j][0]+1)
            j++
        }
    }
    return res
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}