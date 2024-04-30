package longestSubstringWithoutRepeatingCharacters

//leetcode submit region begin(Prohibit modification and deletion)
func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	ans, l, r := 1, 0, 1
	//m := make(map[byte]bool)
	m := make([]int, 128)
	m[s[l]] = 1

	for r < len(s) {
		if m[s[r]] == 1 {
			m[s[l]] = 0
			l++
			continue
		}
		m[s[r]] = 1
		ans = max(ans, r-l+1)
		r++
	}

	return ans

}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}

//leetcode submit region end(Prohibit modification and deletion)
