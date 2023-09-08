package uniquePathsIi

//leetcode submit region begin(Prohibit modification and deletion)
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	dp := make([][]int, len(obstacleGrid))
	for i := 0; i < len(obstacleGrid); i++ {
		dp[i] = make([]int, len(obstacleGrid[i]))
		if obstacleGrid[i][0] == 1 || (i > 0 && dp[i-1][0] == 0) {
			dp[i][0] = 0
		} else {
			dp[i][0] = 1
		}

	}
	for i := 1; i < len(obstacleGrid[0]); i++ {
		if obstacleGrid[0][i] == 1 || (i > 0 && dp[0][i-1] == 0) {
			dp[0][i] = 0
		} else {
			dp[0][i] = 1
		}

	}

	for i := 1; i < len(obstacleGrid); i++ {
		for j := 1; j < len(obstacleGrid[i]); j++ {
			if obstacleGrid[i][j] == 1 {
				dp[i][j] = 0
			} else {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}

	return dp[len(obstacleGrid)-1][len(obstacleGrid[0])-1]
}

//leetcode submit region end(Prohibit modification and deletion)
