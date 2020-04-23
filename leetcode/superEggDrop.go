package leetcode

func superEggDrop(K int, N int) int {
	var dp [][]int
	for i := 0; i <= K; i++ {
		item := make([]int, N+1)
		dp = append(dp, item)
	}
	var m int
	for dp[K][m] < N {
		m++
		for k := 1; k <= K; k++ {
			dp[k][m] = dp[k-1][m-1] + dp[k][m-1] + 1
		}
	}
	return m
}
