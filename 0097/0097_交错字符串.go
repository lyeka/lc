package interleavingString

//leetcode submit region begin(Prohibit modification and deletion)
func isInterleave(s1 string, s2 string, s3 string) bool {
	n, m, l := len(s1), len(s2), len(s3)
	if n+m != l {
		return false
	}

	dp := make([][]bool, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]bool, m+1)
	}

	dp[0][0] = true

	for i := 0; i <= n; i++ {
		for j := 0; j <= m; j++ {
			p := i + j - 1
			if i > 0 {
				dp[i][j] = dp[i][j] || (dp[i-1][j] && s1[i-1] == s3[p])
			}
			if j > 0 {
				dp[i][j] = dp[i][j] || (dp[i][j-1] && s2[j-1] == s3[p])
			}

		}
	}

	return dp[n][m]

}

//leetcode submit region end(Prohibit modification and deletion)
