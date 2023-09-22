package maximalSquare

//leetcode submit region begin(Prohibit modification and deletion)
func maximalSquare(matrix [][]byte) int {
	dp := make([][]int, len(matrix))
	res := 0
	for i := 0; i < len(matrix); i++ {
		dp[i] = make([]int, len(matrix[i]))
		for j := 0; j < len(dp[i]); j++ {
			//if matrix[i][j] == '1' {
			//	dp[i][j] = 1
			//	res = 1
			//}

			dp[i][j] = int(matrix[i][j] - '0')
			if dp[i][j] == 1 {
				res = 1
			}
		}

	}

	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[i]); j++ {
			if matrix[i][j] != '1' {
				continue
			}
			dp[i][j] = min3(dp[i-1][j-1], dp[i-1][j], dp[i][j-1]) + 1
			res = max(res, dp[i][j])
		}
	}

	return res * res

}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min3(x, y, z int) int {
	t := min2(x, y)
	return min2(t, z)
}

func min2(x, y int) int {
	if x < y {
		return x
	}
	return y
}

//leetcode submit region end(Prohibit modification and deletion)
