package longestValidParentheses

//leetcode submit region begin(Prohibit modification and deletion)
func longestValidParentheses(s string) int {
	dp := make([]int, len(s))
	var max int
	for i := 1; i < len(s); i++ {
		if s[i] == ')' {
			if s[i-1] == '(' {
				if i >= 2 {
					dp[i] = dp[i-2] + 2 // e.g. (()()
				} else {
					dp[i] = 2 // e.g. ()
				}
			} else if i-dp[i-1] > 0 && s[i-dp[i-1]-1] == '(' {
				if i-dp[i-1] >= 2 {
					dp[i] = dp[i-1] + 2 + dp[i-dp[i-1]-2] // e.g. ()((()))
				} else {
					dp[i] = dp[i-1] + 2 // e.g. ((()))
				}
			}

			if dp[i] > max {
				max = dp[i]
			}
		}
	}

	return max

}

//leetcode submit region end(Prohibit modification and deletion)
