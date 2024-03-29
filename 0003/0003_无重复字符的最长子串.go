package longestSubstringWithoutRepeatingCharacters

//leetcode submit region begin(Prohibit modification and deletion)
func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	m := make([]int, 128)

	start, end := 0, 1
	res := 1
	m[s[start]]++
	for end < len(s) {
		if m[s[end]] != 1 {
			m[s[end]] = 1
			res = max(res, end-start+1)
			end++
		} else {
			m[s[start]] = 0
			start++

		}
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
