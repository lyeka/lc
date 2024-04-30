package productOfArrayExceptSelf

//leetcode submit region begin(Prohibit modification and deletion)
// 思路：利用前缀和计算上下三角的乘积
func productExceptSelf(nums []int) []int {
	ans := make([]int, len(nums))
	ans[0] = 1

	for i := 1; i < len(nums); i++ {
		ans[i] = ans[i-1] * nums[i-1]
	}

	tmp := 1
	for i := len(nums) - 2; i >= 0; i-- {
		tmp *= nums[i+1]
		ans[i] *= tmp
	}

	return ans

}

//leetcode submit region end(Prohibit modification and deletion)
