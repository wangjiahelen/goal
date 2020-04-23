package main

import "fmt"

func coinChange(coins []int, amount int) int {
	return dp(coins, amount)
}

func dp(coins []int, n int) int {
	if n == 0 {
		return 0
	}
	if n < 0 {
		return -1
	}
	res := int(^uint(0) >> 1)
	for _, v := range coins {
		subproblem := dp(coins, n-v)
		if subproblem == -1 {
			continue
		}
		if subproblem+1 < res {
			res = subproblem + 1
		}
	}
	return res
}

func main() {
	coins := []int{1, 2, 5}
	num := coinChange(coins, 11)
	fmt.Println(num)
}
