package editDistance

//leetcode submit region begin(Prohibit modification and deletion)
func minDistance(word1 string, word2 string) int {
	dp := make([][]int, len(word1)+1)
	for i := 0; i <= len(word1); i++ {
		dp[i] = make([]int, len(word2)+1)
		dp[i][0] = i
	}

	for i := 0; i <= len(word2); i++ {
		dp[0][i] = i
	}

	for i := 1; i <= len(word1); i++ {
		for j := 1; j <= len(word2); j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min3(dp[i-1][j-1], dp[i-1][j], dp[i][j-1]) + 1
			}
		}
	}

	return dp[len(word1)][len(word2)]
}

func min(x, y int) int {
	if x < y {
		return x
	}

	return y
}

func min3(x, y, z int) int {
	return min(min(x, y), z)
}

//leetcode submit region end(Prohibit modification and deletion)
