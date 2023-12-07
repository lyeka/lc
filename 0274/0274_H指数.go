package hIndex

//leetcode submit region begin(Prohibit modification and deletion)
func hIndex(citations []int) int {
	n := len(citations)
	cnt := make([]int, len(citations)+1)
	for _, v := range citations {
		if v >= n {
			cnt[n]++
		} else {
			cnt[v]++
		}
	}
	h := 0
	for i := n; ; i-- {
		h += cnt[i]
		if h >= i {
			return i
		}
	}

}

//leetcode submit region end(Prohibit modification and deletion)
