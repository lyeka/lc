package removeDuplicatesFromSortedArrayIi

//leetcode submit region begin(Prohibit modification and deletion)
func removeDuplicates(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}
	slow, fast := 2, 2
	for fast < len(nums) {
		if nums[slow-2] != nums[fast] {
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}
	return slow
}

//leetcode submit region end(Prohibit modification and deletion)
