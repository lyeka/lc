package twoSum

//leetcode submit region begin(Prohibit modification and deletion)
func twoSum(nums []int, target int) []int {
	m := map[int]int{}
	for i, v := range nums {
		i2, ok := m[target-v]
		if ok {
			return []int{i, i2}
		}
		m[v] = i
	}

	return nil
}

//leetcode submit region end(Prohibit modification and deletion)
