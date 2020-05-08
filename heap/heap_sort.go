package heap

func buildHeap(a []int, n int) {
	//heapify from the last parent node
	for i := n / 2; i >= 1; i-- {
		heapifyUpToDown(a, i, n)
	}
}

//sort by ascend, a index begin from 1, has n elements
func sort(a []int, n int) {
	buildHeap(a, n)

	k := n
	for k >= 1 {
		a[1], a[k] = a[k], a[1]
		heapifyUpToDown(a, 1, k-1)
		k--
	}
}

//heapify from up to down, node index = top
func heapifyUpToDown(a []int, top, count int) {
	for i := top; i <= count/2; {
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
