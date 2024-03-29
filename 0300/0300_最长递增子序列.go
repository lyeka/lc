package longestIncreasingSubsequence

//leetcode submit region begin(Prohibit modification and deletion)
func lengthOfLIS(nums []int) int {
	dp := make([]int, len(nums))

	res := 0

	for i := 0; i < len(nums); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}

		res = max(res, dp[i])
	}

	return res
}
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

//leetcode submit region end(Prohibit modification and deletion)
