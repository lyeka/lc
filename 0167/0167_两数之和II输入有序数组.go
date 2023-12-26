package twoSumIiInputArrayIsSorted

//leetcode submit region begin(Prohibit modification and deletion)
func twoSum(numbers []int, target int) []int {
	l, r := 0, len(numbers)-1
	for l < r {
		t := numbers[l] + numbers[r]
		if t == target {
			return []int{l + 1, r + 1}
		}
		if t < target {
			l++
		} else {
			r--
		}
	}
	return nil
}

//leetcode submit region end(Prohibit modification and deletion)
