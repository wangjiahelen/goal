package sort2

//归并排序，非原地排序，因为需要合并，由下到上，先处理子问题，后合并
func MergeSort(arr []int) {
	arrLen := len(arr)
	if arrLen <= 1 {
		return
	}
	mergeSort(arr, 0, arrLen-1)
}

func mergeSort(arr []int, start, end int) {
	if start >= end {
		return
	}
	mid := start + (end-start)/2
	mergeSort(arr, start, mid)
	mergeSort(arr, mid+1, end)
	merge(arr, start, mid, end)
}

func merge(arr []int, start, mid, end int) {
	tmpArr := make([]int, end-start+1)
	i := start
	j := mid + 1
	k := 0
	for ; i <= mid && j <= end; k++ {
		if arr[i] <= arr[j] {
			tmpArr[k] = arr[i]
			i++
		} else {
			tmpArr[k] = arr[j]
			j++
		}
	}
	for i <= mid {
		tmpArr[k] = arr[i]
		i++
		k++
	}

	for j <= end {
		tmpArr[k] = arr[j]
		j++
		k++
	}
	copy(arr[start:end+1], tmpArr)
}

//快速排序是原地排序，从上往下，先分区，后处理子问题。原地分区
func QuickSort(arr []int) {
	separateSort(arr, 0, len(arr)-1)
}

func separateSort(arr []int, start, end int) {
	if start >= end {
		return
	}
	i := partition(arr, start, end)
	separateSort(arr, start, i-1)
	separateSort(arr, i+1, end)
}

func partition(arr []int, start, end int) int {
	//选取最后一位作为比较数字
	pivot := arr[end]

	i := start
	for j := start; j < end; j++ {
		if arr[j] < pivot {
			if !(i == j) {
				//swap
				arr[i], arr[j] = arr[j], arr[i]
			}
			i++
		}
	}
	arr[i], arr[end] = arr[end], arr[i]
	return i
}

//编程实现 O(n) 时间复杂度内找到一组数据的第 K 大元素，使用快速排序中分区，如果i等于K则返回，如果i>K，则在左边查找，如果i<K，则在右边查找
func FindKthLargest(arr []int, k int) int {
	if len(arr) == 0 || k <= 0 || k > len(arr) {
		return -1
	}

	return quickSelect(arr, 0, len(arr)-1, k)
}

func quickSelect(arr []int, start, end, k int) int {
	if start == end {
		return arr[start]
	}
	left, right, mid := start, end, start+(end-start)/2
	pivot := arr[mid]
	for left <= right {
		for left <= right && arr[left] > pivot {
			left++
		}
		for left <= right && arr[right] < pivot {
			right--
		}
		if left <= right {
			arr[left], arr[right] = arr[right], arr[left]
			left++
			right--
		}
	}
	if start+k-1 <= right {
		return quickSelect(arr, start, right, k)
	}
	if start+k-1 >= left {
		return quickSelect(arr, left, end, k-(left-start))
	}
	return arr[right+1]
}
