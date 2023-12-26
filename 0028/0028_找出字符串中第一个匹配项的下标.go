package findTheIndexOfTheFirstOccurrenceInAString

//leetcode submit region begin(Prohibit modification and deletion)
func strStr(haystack string, needle string) int {
	for i := 0; i < len(haystack); i++ {
		end := i + len(needle)
		if end > len(haystack) {
			return -1
		}
		if haystack[i:i+len(needle)] == needle {
			return i
		}
	}
	return -1

}

//leetcode submit region end(Prohibit modification and deletion)
