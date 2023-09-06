package minimumPathSum

//leetcode submit region begin(Prohibit modification and deletion)
func minPathSum(grid [][]int) int {
	dp := make([][]int, len(grid))
	for i := 0; i < len(grid); i++ {
		dp[i] = make([]int, len(grid[i]))
		if i >= 1 {
			dp[i][0] = grid[i][0] + dp[i-1][0]
		} else {
			dp[i][0] = grid[i][0]
		}

	}
	for j := 1; j < len(grid[0]); j++ {
		dp[0][j] = grid[0][j] + dp[0][j-1]
	}

	for i := 1; i < len(grid); i++ {
		for j := 1; j < len(grid[i]); j++ {
			dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
		}
	}

	return dp[len(grid)-1][len(grid[len(grid)-1])-1]

}

func min(x, y int) int {
	if x < y {
		return x
	}

	return y
}

//leetcode submit region end(Prohibit modification and deletion)
