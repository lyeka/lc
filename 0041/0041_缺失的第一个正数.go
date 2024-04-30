package firstMissingPositive

//leetcode submit region begin(Prohibit modification and deletion)
func firstMissingPositive(nums []int) int {
	n := len(nums)
	for i := 0; i < n; i++ { // 1. 移除非正数
		if nums[i] <= 0 {
			nums[i] = n + 1
		}
	}

	for i := 0; i < n; i++ { // 2. 位置打标
		x := abs(nums[i])
		if x > 0 && x <= n {
			nums[x-1] = -abs(nums[x-1])
		}
	}

	for i := 0; i < n; i++ {
		if nums[i] > 0 {
			return i + 1
		}
	}

	return n + 1

}

func abs(i int) int {
	if i < 0 {
		return -i
	}

	return i
}

//leetcode submit region end(Prohibit modification and deletion)
