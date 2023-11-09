package majorityElement

//leetcode submit region begin(Prohibit modification and deletion)
func majorityElement(nums []int) int {
	m := make(map[int]int)
	for _, v := range nums {
		_, ok := m[v]
		if ok {
			m[v]++
		} else {
			m[v] = 1
		}

		if m[v] > len(nums)/2 {
			return v
		}

	}

	return 0

}

//leetcode submit region end(Prohibit modification and deletion)
