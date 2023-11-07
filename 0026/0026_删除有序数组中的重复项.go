package removeDuplicatesFromSortedArray

//leetcode submit region begin(Prohibit modification and deletion)
func removeDuplicates(nums []int) int {
	left := 0
	for _, v := range nums {
		if nums[left] == v {
			continue
		}
		left++
		nums[left] = v
	}

	return left + 1

}

//leetcode submit region end(Prohibit modification and deletion)
