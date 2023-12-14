package trappingRainWater

//leetcode submit region begin(Prohibit modification and deletion)
func trap(height []int) int {
	left, right, leftMax, rightMax := 0, len(height)-1, height[0], height[len(height)-1]
	res := 0
	for left < right {
		leftMax = max(height[left], leftMax)
		rightMax = max(height[right], rightMax)
		if height[left] < height[right] {
			res += leftMax - height[left]
			left++
		} else {
			res += rightMax - height[right]
			right--
		}
	}

	return res
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
