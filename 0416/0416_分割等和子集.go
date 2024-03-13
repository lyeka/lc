package partitionEqualSubsetSum

//leetcode submit region begin(Prohibit modification and deletion)
func canPartition(nums []int) bool {
	if len(nums) < 2 {
		return false
	}
	total := 0
	for _, num := range nums {
		total += num
	}

	if total%2 != 0 {
		return false
	}

	target := total / 2

	dp := make([][]bool, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = make([]bool, target+1)
		dp[i][0] = true
	}

	for i := 1; i < len(nums); i++ {
		for j := 1; j <= target; j++ {
			if j >= nums[i] {
				dp[i][j] = dp[i-1][j] || dp[i-1][j-nums[i]]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}

	return dp[len(nums)-1][target]
}

//leetcode submit region end(Prohibit modification and deletion)
