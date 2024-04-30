package trappingRainWater

//leetcode submit region begin(Prohibit modification and deletion)
func trap(height []int) int {
	ans := 0
	left, right, leftMax, rightMax := 0, len(height)-1, height[0], height[len(height)-1]
	for left < right {
		leftMax = max(leftMax, height[left])
		rightMax = max(rightMax, height[right])
		if height[left] < height[right] {
			ans += leftMax - height[left]
			left++
		} else {
			ans += rightMax - height[right]
			right--
		}
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
