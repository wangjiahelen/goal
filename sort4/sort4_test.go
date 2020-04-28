package sort4

import (
	"fmt"
	"testing"
)

func TestCountingSort(t *testing.T) {
	arr := []int{5, 4}
	CountingSort(arr, len(arr))
	fmt.Println(arr)

	arr = []int{5, 4, 3, 2, 1}
	CountingSort(arr, len(arr))
	fmt.Println(arr)
}
