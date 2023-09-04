package longestSubstringWithoutRepeatingCharacters

//leetcode submit region begin(Prohibit modification and deletion)
func lengthOfLongestSubstring(s string) int {
	var m [128]int
	//m := make([]int, 128)
	i := 0
	res := 0
	for j := 0; j < len(s); j++ {
		i = max(i, m[s[j]])
		m[s[j]] = j + 1
		res = max(res, j-i+1)
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
