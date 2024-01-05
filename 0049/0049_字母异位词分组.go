package groupAnagrams

//leetcode submit region begin(Prohibit modification and deletion)
func groupAnagrams(strs []string) [][]string {
	m := map[[26]int][]string{}
	for _, str := range strs {
		cnt := [26]int{}
		for _, b := range str {
			cnt[b-'a']++
		}
		m[cnt] = append(m[cnt], str)
	}

	res := [][]string{}
	for _, v := range m {
		res = append(res, v)
	}

	return res

}

//leetcode submit region end(Prohibit modification and deletion)
