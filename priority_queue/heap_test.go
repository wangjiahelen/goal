package pqueue

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_AdjustHeap(t *testing.T) {
	list := []Node{{0, 0}, {1, 1}, {2, 2}, {3, 3}, {4, 1}, {6, 6}}

	adjustHeap(list, 1, len(list)-1)
	fmt.Println(list[1].value, list[1].priority)
	assert.Equal(t, 6, list[1].value)
}
