package pqueue

type Node struct {
	value    int
	priority int
}

type PQueue struct {
	heap     []Node
	capacity int
	used     int
}

func NewPriorityQueue(capacity int) PQueue {
	return PQueue{
		heap:     make([]Node, capacity+1, capacity+1),
		capacity: capacity,
		used:     0,
	}
}

func (q *PQueue) Push(node Node) {
	if q.used > q.capacity {
		//队列已满
		return
	}
	q.used++
	q.heap[q.used] = node
	//堆化可以放在pop中
	//adjustHeap(q.heap, 1, q.used)
}

func (q *PQueue) Pop() Node {
	if q.used == 0 {
		return Node{-1, -1}
	}
	//先堆化，再取堆顶元素
	adjustHeap(q.heap, 1, q.used)
	node := q.heap[1]

	q.heap[1] = q.heap[q.used]
	q.used--
	return node
}

func (q *PQueue) Top() Node {
	if q.used == 0 {
		return Node{-1, -1}
	}
	adjustHeap(q.heap, 1, q.used)
	return q.heap[1]
}
