package dp

import "fmt"

func lsc(s1, s2 string) int {
	m, n := len(s1), len(s2)

	memo := make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		memo[i] = make([]int, n+1)
	}

	for i := 1; i < m+1; i++ {
		for j := 1; j < n+1; j++ {
			if s1[i-1] == s2[j-1] {
				memo[i][j] = memo[i-1][j-1] + 1
			}
		}
	}
	fmt.Println(memo)
	longest := 0
	for i := range memo {
		for j, e2 := range memo[i] {
			if longest < memo[i][j] {
				longest = e2
			}
		}
	}
	return longest
}
