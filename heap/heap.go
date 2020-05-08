package heap

type Heap struct {
	a     []int
	n     int
	count int
}

func NewHeap(capacity int) *Heap {
	heap := &Heap{}
	heap.n = capacity
	heap.a = make([]int, capacity+1)
	heap.count = 0
	return heap
}

//top-max heap -> heapify from down to up
func (heap *Heap) insert(data int) {
	//defensive
	if heap.count == heap.n {
		return
	}
	heap.count++
	heap.a[heap.count] = data

	//compare with parent node
	i := heap.count
	parent := i / 2
	for parent > 0 && heap.a[parent] < heap.a[i] {
		heap.a[parent], heap.a[i] = heap.a[i], heap.a[parent]
		i = parent
		parent = i / 2
	}
}

//heapfify from up to down
func (heap *Heap) removeMax() {
	//defensive
	if heap.count == 0 {
		return
	}
	heap.a[1], heap.a[heap.count] = heap.a[heap.count], heap.a[1]
	heap.count--

	heapfifyUpToDown(heap.a, heap.count)
}

func heapfifyUpToDown(a []int, count int) {
	for i := 1; i <= count/2; {
		maxIndex := i
		if a[i] < a[i*2] {
			maxIndex = i * 2
		}

		if i*2+1 <= count && a[maxIndex] < a[i*2+1] {
			maxIndex = i*2 + 1
		}

		if maxIndex == i {
			break
		}
		a[i], a[maxIndex] = a[maxIndex], a[i]
		i = maxIndex
	}
}
