package longestCommonPrefix

//leetcode submit region begin(Prohibit modification and deletion)
func longestCommonPrefix(strs []string) string {
	if len(strs) == 1 {
		return strs[0]
	}

	pre := make([]byte, 0)
	for i, _ := range strs[0] {
		for j := 1; j < len(strs); j++ {
			if i >= len(strs[j]) || strs[j][i] != strs[0][i] {
				return string(pre)
			}
		}
		pre = append(pre, strs[0][i])
	}

	return string(pre)

}

//leetcode submit region end(Prohibit modification and deletion)
