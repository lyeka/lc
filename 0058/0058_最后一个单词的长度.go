package lengthOfLastWord

//leetcode submit region begin(Prohibit modification and deletion)
func lengthOfLastWord(s string) int {
	start, end := 0, -1
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' && end == -1 {
			continue
		}
		if s[i] != ' ' && end == -1 {
			end = i
		}

		if s[i] == ' ' {
			start = i + 1
			break
		}

	}

	return end - start + 1

}

//leetcode submit region end(Prohibit modification and deletion)
