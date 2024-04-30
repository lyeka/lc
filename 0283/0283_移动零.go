package moveZeroes

//leetcode submit region begin(Prohibit modification and deletion)
func moveZeroes(nums []int) {
	l, r, n := 0, 0, len(nums)
	for r < n {
		if nums[r] != 0 {
			nums[l], nums[r] = nums[r], nums[l]
			l++
		}
		r++
	}
}

//leetcode submit region end(Prohibit modification and deletion)
