package maximumProductSubarray

import (
	"math"
)

//leetcode submit region begin(Prohibit modification and deletion)
func maxProduct(nums []int) int {
	dpMin := make([]int, len(nums))
	dpMax := make([]int, len(nums))
	ans := math.MinInt
	for i := 0; i < len(nums); i++ {
		dpMax[i], dpMin[i] = nums[i], nums[i]
		if i > 0 {
			dpMax[i] = max(dpMax[i-1]*nums[i], max(nums[i], dpMin[i-1]*nums[i]))
			dpMin[i] = min(dpMin[i-1]*nums[i], min(nums[i], dpMax[i-1]*nums[i]))
		}
		ans = max(ans, dpMax[i])
	}

	return ans

}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}

	return y
}

//leetcode submit region end(Prohibit modification and deletion)
