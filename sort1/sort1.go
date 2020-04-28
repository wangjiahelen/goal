package sort1

/*冒泡排序、插入排序、选择排序*/
func BubbleSort(a []int, n int) {
	if n <= 1 {
		return
	}
	for i := 1; i < n; i++ {
		flag := false
		for j := 0; j < n-i-1; j++ {
			if a[j] > a[j+1] {
				a[j+1], a[j] = a[j], a[j+1]
				flag = true
			}
		}
		if !flag {
			break
		}
	}
}

func InsertionSort(a []int, n int) {
	if n <= 1 {
		return
	}
	for i := 1; i < n; i++ {
		value := a[i]
		j := i - 1
		for ; j >= 0; j-- {
			if a[j] > value {
				a[j+1] = a[j]
			} else {
				break
			}
		}
		a[j+1] = value
	}
}

func SelectionSort(a []int, n int) {
	if n <= 1 {
		return
	}
	for i := 0; i < n; i++ {
		minIndex := i
		for j := i + 1; j < n; j++ {
			if a[j] < a[minIndex] {
				minIndex = j
			}
		}
		//交换
		a[i], a[minIndex] = a[minIndex], a[i]
	}
}
