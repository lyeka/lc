package minimumSizeSubarraySum

//leetcode submit region begin(Prohibit modification and deletion)
func minSubArrayLen(target int, nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	res := len(nums) + 1
	start, end := 0, 0
	total := 0
	for end < len(nums) {
		total += nums[end]
		for total >= target {
			res = min(res, end-start+1)
			total -= nums[start]
			start++
		}
		end++

	}

	if res == len(nums)+1 {
		return 0
	} else {
		return res
	}

}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

//leetcode submit region end(Prohibit modification and deletion)
