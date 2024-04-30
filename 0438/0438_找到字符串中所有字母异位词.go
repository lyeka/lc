package findAllAnagramsInAString

//leetcode submit region begin(Prohibit modification and deletion)
func findAnagrams(s string, p string) []int {
	if len(s) < len(p) {
		return nil
	}
	pk, sk := [26]int{}, [26]int{}
	n := len(p)
	for i := 0; i < n; i++ {
		pk[p[i]-'a']++
		sk[s[i]-'a']++
	}

	ans := []int{}
	l, r := 0, n-1
	for r < len(s) {
		if pk == sk {
			ans = append(ans, l)
		}
		sk[s[l]-'a']--
		l++
		r++
		if r < len(s) {
			sk[s[r]-'a']++
		}

	}

	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
