package romanToInteger

//leetcode submit region begin(Prohibit modification and deletion)
func romanToInt(s string) int {
	vm := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	res := 0
	for i := 0; i < len(s); i++ {
		if i+1 < len(s) && vm[s[i+1]] > vm[s[i]] {
			res -= vm[s[i]]
		} else {
			res += vm[s[i]]
		}
	}

	return res

}

//leetcode submit region end(Prohibit modification and deletion)
