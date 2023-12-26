package isSubsequence

//leetcode submit region begin(Prohibit modification and deletion)
func isSubsequence(s string, t string) bool {
	ps, pt := 0, 0
	for ps < len(s) && pt < len(t) {
		if s[ps] == t[pt] {
			ps++
		}
		pt++
	}

	return ps == len(s)

}

//leetcode submit region end(Prohibit modification and deletion)
