package jumpGameIi

//leetcode submit region begin(Prohibit modification and deletion)
// 贪心算法：局部最优则全局最优
func jump(nums []int) int {
	end := 0
	rightMost := 0
	steps := 0
	for i := 0; i < len(nums)-1; i++ {
		rightMost = max(rightMost, i+nums[i])
		if i == end {
			end = rightMost
			steps++
		}
	}

	return steps

}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

//leetcode submit region end(Prohibit modification and deletion)
