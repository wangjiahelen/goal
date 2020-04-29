package stringmatch

import (
	"fmt"
	"testing"
)

func Test_bfSearch(t *testing.T) {
	main := "abcd227fac"
	pattern := "ac"
	fmt.Println(bfSearch(main, pattern))
}
