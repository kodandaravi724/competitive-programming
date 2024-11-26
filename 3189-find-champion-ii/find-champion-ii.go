func dfs(visited []bool, adjacencyList map[int][]int, current int) {
    if visited[current] {
        return
    }
    visited[current] = true
    neighbours := adjacencyList[current]
    for _, n := range(neighbours) {
        dfs(visited, adjacencyList, n)
    }
}

func findChampion(n int, edges [][]int) int {
    var adjacencyList map[int][]int = make(map[int][]int)
    var isStronger []bool = []bool{}
    i := 0
    for i < n {
        adjacencyList[i] = []int{}
        isStronger = append(isStronger, false)
        i++
    }
    for _, edge := range(edges) {
        adjacencyList[edge[0]] = append(adjacencyList[edge[0]], edge[1])
    }
    i = 0
    for i < n {
        if !isStronger[i] {
            for _, ne := range(adjacencyList[i]) {
                dfs(isStronger, adjacencyList, ne)
            }
        }
        i++
    }
    numberOfChampions := 0
    i = 0
    ind := 0
    for i < n {
        if !isStronger[i] {
            numberOfChampions++
            ind = i
        }
        i++
    }
    if numberOfChampions == 1 {
        return ind
    }
    return -1
}