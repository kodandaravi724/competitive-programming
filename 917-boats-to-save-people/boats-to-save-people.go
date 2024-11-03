import "sort"

func numRescueBoats(people []int, limit int) int {
    sort.Ints(people)
    var i, j int = 0, len(people) - 1
    var boats = 0
    for (i <= j) {
        weight := people[i] + people[j]
        if weight <= limit {
            i++
            j--
        } else {
            j--
        }
        boats++
    }
    return boats
}