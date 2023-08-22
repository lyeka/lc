package coinChange

//leetcode submit region begin(Prohibit modification and deletion)
func coinChange(coins []int, amount int) int {
	max := amount + 1
	dp := make([]int, amount+1)
	dp[0] = 0
	for i := 1; i <= amount; i++ {
		dp[i] = max
	}

	for i := 1; i <= amount; i++ {
		for _, coin := range coins {
			if i-coin >= 0 {
				dp[i] = min(dp[i-coin]+1, dp[i])
			}
		}
	}

	if dp[amount] == max {
		return -1 // 代表无法组合而成
	}
	return dp[amount]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

//leetcode submit region end(Prohibit modification and deletion)
