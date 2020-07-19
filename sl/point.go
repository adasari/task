package sl

// Point .
type Point struct {
	flag     bool // true if the point is start
	priority int
	NamedRect
}

// A PriorityQueue implements heap.Interface and holds Points.
type PriorityQueue []Point

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

// Push .
func (pq *PriorityQueue) Push(item interface{}) {
	*pq = append(*pq, item.(Point))
}

// Pop .
func (pq *PriorityQueue) Pop() interface{} {
	n := len(*pq)
	item := (*pq)[n-1]
	*pq = (*pq)[:n-1]
	return item
}
