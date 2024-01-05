package ransomNote

//leetcode submit region begin(Prohibit modification and deletion)
func canConstruct(ransomNote string, magazine string) bool {
	rm, mm := [58]int{}, [58]int{}
	for _, v := range ransomNote {
		rm[v-'A']++
	}
	for _, v := range magazine {
		mm[v-'A']++
	}

	for i, rv := range rm {
		mv := mm[i]
		if rv > mv {
			return false
		}
	}

	return true
}

//leetcode submit region end(Prohibit modification and deletion)
