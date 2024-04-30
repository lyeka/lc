package maximumSubarray

//leetcode submit region begin(Prohibit modification and deletion)
func maxSubArray(nums []int) int {
	ans := nums[0]
	for i := 1; i < len(nums); i++ {
		nums[i] = max(nums[i-1]+nums[i], nums[i])
		ans = max(ans, nums[i])
	}
	return ans

}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}

//leetcode submit region end(Prohibit modification and deletion)
