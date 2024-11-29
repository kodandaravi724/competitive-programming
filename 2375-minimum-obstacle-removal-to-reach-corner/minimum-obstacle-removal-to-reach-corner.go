
var directions = [][2]int{
	{-1, 0}, {1, 0}, {0, -1}, {0, 1},
}

type Cell struct {
	x, y, obstaclesRemoved int
}

type MinHeap []Cell

func (h MinHeap) Len() int            { return len(h) }
func (h MinHeap) Less(i, j int) bool  { return h[i].obstaclesRemoved < h[j].obstaclesRemoved }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.(Cell)) }
func (h *MinHeap) Pop() interface{}   { 
	n := len(*h)
	item := (*h)[n-1]
	*h = (*h)[:n-1]
	return item
}

func minimumObstacles(grid [][]int) int {
	m, n := len(grid), len(grid[0])

	pq := &MinHeap{{x: 0, y: 0, obstaclesRemoved: 0}}
	heap.Init(pq)

	minRemoved := make([][]int, m)
	for i := 0; i < m; i++ {
		minRemoved[i] = make([]int, n)
		for j := 0; j < n; j++ {
			minRemoved[i][j] = math.MaxInt
		}
	}
	minRemoved[0][0] = 0

	for pq.Len() > 0 {
		cell := heap.Pop(pq).(Cell)

		if cell.x == m-1 && cell.y == n-1 {
			return cell.obstaclesRemoved
		}

		for _, dir := range directions {
			newX, newY := cell.x+dir[0], cell.y+dir[1]

			if newX >= 0 && newX < m && newY >= 0 && newY < n {
				newObstaclesRemoved := cell.obstaclesRemoved + grid[newX][newY]

				if newObstaclesRemoved < minRemoved[newX][newY] {
					minRemoved[newX][newY] = newObstaclesRemoved
					heap.Push(pq, Cell{x: newX, y: newY, obstaclesRemoved: newObstaclesRemoved})
				}
			}
		}
	}

	return -1
}