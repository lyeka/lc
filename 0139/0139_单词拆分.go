package wordBreak

//leetcode submit region begin(Prohibit modification and deletion)
func wordBreak(s string, wordDict []string) bool {
	wordSet := make(map[string]bool)
	for _, word := range wordDict {
		wordSet[word] = true
	}

	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if dp[j] && wordSet[s[j:i]] {
				dp[i] = true
			}
		}
	}

	return dp[len(s)]
}

//leetcode submit region end(Prohibit modification and deletion)
