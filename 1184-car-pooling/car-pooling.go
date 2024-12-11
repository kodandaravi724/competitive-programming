func carPooling(trips [][]int, capacity int) bool {
    var numPassengers []int = make([]int, 1001)

    for _, trip := range(trips) {
        numPassengers[trip[1]] += trip[0]
        numPassengers[trip[2]] -= trip[0]
    }

    numP := 0

    for _, p := range(numPassengers) {
        numP += p
        if numP > capacity {
            return false
        }
    }
    return true
}