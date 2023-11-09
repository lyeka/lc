package rotateArray

//leetcode submit region begin(Prohibit modification and deletion)
func rotate(nums []int, k int) {
	reverse(nums)
	k = k % len(nums)
	reverse(nums[0:k])
	reverse(nums[k:])
}

func reverse(nums []int) {
	n := len(nums)
	for i := 0; i < n/2; i++ {
		nums[i], nums[n-i-1] = nums[n-i-1], nums[i]
	}
}

//leetcode submit region end(Prohibit modification and deletion)
