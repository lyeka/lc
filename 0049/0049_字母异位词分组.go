package groupAnagrams

//leetcode submit region begin(Prohibit modification and deletion)
func groupAnagrams(strs []string) [][]string {
	m := make(map[[26]int][]string)
	for _, str := range strs {
		key := [26]int{}
		for i := 0; i < len(str); i++ {
			key[str[i]-'a']++
		}
		m[key] = append(m[key], str)
	}

	ans := make([][]string, 0)
	for _, v := range m {
		ans = append(ans, v)
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
