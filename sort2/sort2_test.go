package sort2

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestMergeSort(t *testing.T) {
	arr := []int{5, 4}
	MergeSort(arr)
	fmt.Println(arr)

	arr = []int{5, 4, 3, 2, 1}
	MergeSort(arr)
	fmt.Println(arr)
}

func TestQuickSort(t *testing.T) {
	arr := []int{5, 4}
	QuickSort(arr)
	fmt.Println(arr)

	arr = createRandomArr(100)
	QuickSort(arr)
	fmt.Println(arr)
}

func createRandomArr(length int) []int {
	arr := make([]int, length, length)
	for i := 0; i < length; i++ {
		arr[i] = rand.Intn(100)
	}
	return arr
}

func TestFindKthLargest(t *testing.T) {
	arr := []int{5, 4, 3, 2, 1}
	k := FindKthLargest(arr, 3)
	fmt.Println(k)
	arr = []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	k = FindKthLargest(arr, 10)
	fmt.Println(k)
}
