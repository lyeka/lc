package removeElement

//leetcode submit region begin(Prohibit modification and deletion)
func removeElement(nums []int, val int) int {
	left := 0
	for _, v := range nums {
		if v != val {
			nums[left] = v
			left++
		}

	}

	return left

}

//leetcode submit region end(Prohibit modification and deletion)
