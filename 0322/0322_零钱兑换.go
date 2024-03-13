package coinChange

//leetcode submit region begin(Prohibit modification and deletion)
func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	max := amount + 1
	for i := 1; i <= amount; i++ {
		dp[i] = max
		for _, v := range coins {
			if i-v < 0 {
				continue
			}
			dp[i] = min(dp[i], dp[i-v]+1)
		}

	}

	if dp[amount] == max {
		return -1
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
