package dp

import (
	"fmt"
	"testing"
)

func Test_lcs(t *testing.T) {
	fmt.Println(lsc("blue", "clues")) //3
	fmt.Println(lsc("fosh", "fish"))  //2
	fmt.Println(lsc("fosh", "fort"))  //2
	fmt.Println(lsc("hish", "fish"))  //3
	fmt.Println(lsc("hish", "vista")) //2
}
