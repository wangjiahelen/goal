package bitmap

import (
	"fmt"
	"testing"
)

func TestBitMap(t *testing.T) {
	bitMap := New(80)
	for i := 0; i <= 100; i += 10 {
		bitMap.Set(uint(i))
	}
	for i := 0; i <= 100; i += 10 {
		fmt.Println(bitMap.Get(uint(i)))
	}
}
