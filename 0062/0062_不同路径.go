package uniquePaths

//leetcode submit region begin(Prohibit modification and deletion)
func uniquePaths(m int, n int) int {
	dp := make([][]int, m+1)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n+1)
		dp[i][0] = 1
	}

	for i := 0; i < n; i++ {
		dp[0][i] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}

	return dp[m-1][n-1]

}

//leetcode submit region end(Prohibit modification and deletion)
