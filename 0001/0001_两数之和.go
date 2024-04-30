package twoSum

//leetcode submit region begin(Prohibit modification and deletion)
func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if index, ok := m[target-nums[i]]; ok {
			return []int{index, i}
		}
		m[nums[i]] = i
	}

	return nil
}

//leetcode submit region end(Prohibit modification and deletion)
