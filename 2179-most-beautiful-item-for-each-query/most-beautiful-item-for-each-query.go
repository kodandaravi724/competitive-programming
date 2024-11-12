import "sort"

type ElementInd struct {
    element int
    index int
}

func max(i, j int) int {
    if i > j {
        return i
    }
    return j
}

func maximumBeauty(items [][]int, queries []int) []int {
    var queriesInd []ElementInd = []ElementInd{}
    var result []int = []int{}
    for i,_ := range(queries) {
        result = append(result, i)
    }
    for index, query := range(queries) {
        queriesInd = append(queriesInd, ElementInd{
            query, index, 
        })
    }
    sort.Slice(queriesInd, func (i, j int) bool {
        return queriesInd[i].element < queriesInd[j].element
    })
    sort.Slice(items, func (i, j int) bool {
        return items[i][0] < items[j][0]
    })
    maxValue := 0
    i := 0
    j := 0
    for j < len(queriesInd) {
        for i < len(items) && (queriesInd[j].element >= items[i][0]) {
                maxValue = max(maxValue, items[i][1])
                i++
        }
        result[queriesInd[j].index] = maxValue
        j++
    }
    return result
}