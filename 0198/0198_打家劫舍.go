package houseRobber

//leetcode submit region begin(Prohibit modification and deletion)
func rob(nums []int) int {
	p, q, r := 0, 0, 0
	for i := 0; i < len(nums); i++ {
		p = q
		q = r
		r = max(p+nums[i], q)
	}

	return r
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

//leetcode submit region end(Prohibit modification and deletion)
