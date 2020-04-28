package sort3

import (
	"fmt"
	"testing"
)

func TestBucketSort(t *testing.T) {
	a := []int{1, 6, 3, 5, 8, 6, 4}
	BucketSort(a)
	fmt.Println(a)
}

func TestBucketSortSimple(t *testing.T) {
	a := []int{1, 6, 3, 5, 8, 6, 4}
	BucketSortSimple(a)
	fmt.Println(a)
}
