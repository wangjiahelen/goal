package dp

import (
	"fmt"
	"testing"
)

func TestFindLeastCoins(t *testing.T) {

	coinOptions := []int{1, 3, 5, 10, 50}

	Cnt = 0
	result := LeastCoins(9, coinOptions)

	fmt.Println("test 1 =====================")
	if result != 3 {
		fmt.Printf("least coins %d", result)
		t.Error("failed")
	}
	fmt.Printf("cnt===%d\n", Cnt)

	Cnt = 0
	fmt.Println("test 2 =====================")
	result = LeastCoins(36, coinOptions)

	if result != 5 {
		fmt.Printf("least coins %d", result)
		t.Error("failed")
	}
	fmt.Printf("cnt===%d\n", Cnt)

}

func TestFindLeastCoins2(t *testing.T) {

	coinOptions := []int{1, 3, 5, 10, 50}

	Cnt = 0
	result := LeastCoins2(9, coinOptions)

	fmt.Println("test 1 =====================")
	if result != 3 {
		fmt.Printf("least coins %d", result)
		t.Error("failed")
	}

	fmt.Printf("cnt===%d\n", Cnt)

	Cnt = 0
	fmt.Println("test 2 =====================")
	result = LeastCoins2(36, coinOptions)

	if result != 5 {
		fmt.Printf("least coins %d", result)
		t.Error("failed")
	}
	fmt.Printf("cnt===%d\n", Cnt)

}
