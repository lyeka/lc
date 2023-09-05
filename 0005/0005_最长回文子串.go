package longestPalindromicSubstring

//leetcode submit region begin(Prohibit modification and deletion)
// 中心扩展
func longestPalindrome(s string) string {
	start := 0
	maxLen := 1
	for i := 0; i < len(s); i++ {
		// 对应 aba 类型
		l, r := foo(s, i, i)
		if r-l+1 > maxLen {
			start = l
			maxLen = r - l + 1
		}

		// 对应 abba 类型
		l, r = foo(s, i, i+1)
		if r-l+1 > maxLen {
			start = l
			maxLen = r - l + 1
		}
	}

	return s[start : start+maxLen]

}

func foo(s string, l, r int) (int, int) {
	for l >= 0 && r < len(s) && s[l] == s[r] {
		l--
		r++
	}
	return l + 1, r - 1
}

// dp
//func longestPalindrome(s string) string {
//	start := 0
//	maxLen := 1
//	dp := make([][]bool, len(s))
//	for i := 0; i < len(s); i++ {
//		dp[i] = make([]bool, len(s))
//		dp[i][i] = true
//	}
//
//	for l := 2; l <= len(s); l++ {
//		for i := 0; i < len(s); i++ {
//			j := i + l - 1
//
//			if j >= len(s) {
//				break
//			}
//
//			if s[i] != s[j] {
//				dp[i][j] = false
//			} else {
//				if l <= 2 {
//					dp[i][j] = true
//				} else {
//					dp[i][j] = dp[i+1][j-1]
//				}
//
//			}
//
//			if dp[i][j] && l > maxLen {
//				maxLen = l
//				start = i
//			}
//		}
//	}
//
//	return s[start : start+maxLen]
//}

//leetcode submit region end(Prohibit modification and deletion)
