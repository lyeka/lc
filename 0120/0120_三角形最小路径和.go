package triangle

//leetcode submit region begin(Prohibit modification and deletion)
// 时间复杂度：O(n)
// 空间复杂度: O(n^2) // todo 优化空间复杂度
func minimumTotal(triangle [][]int) int {
	dp := make([][]int, len(triangle))
	for i := 0; i < len(triangle); i++ {
		dp[i] = make([]int, len(triangle[i]))
	}
	dp[0][0] = triangle[0][0]
	for i := 1; i < len(triangle); i++ {
		dp[i][0] = dp[i-1][0] + triangle[i][0]
		for j := 1; j < i; j++ {
			dp[i][j] = min(dp[i-1][j], dp[i-1][j-1]) + triangle[i][j]
		}
		dp[i][i] = dp[i-1][i-1] + triangle[i][i]
	}

	res := dp[len(triangle)-1][0]
	for i := 1; i < len(dp[len(triangle)-1]); i++ {
		res = min(res, dp[len(triangle)-1][i])
	}

	return res

}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

//leetcode submit region end(Prohibit modification and deletion)
