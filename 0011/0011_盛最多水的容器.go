package containerWithMostWater

//leetcode submit region begin(Prohibit modification and deletion)
func maxArea(height []int) int {
	l, r, ans := 0, len(height)-1, 0
	for l < r {
		h := min(height[l], height[r])
		area := h * (r - l)
		if area > ans {
			ans = area
		}
		if height[l] < height[r] {
			l++
		} else {
			r--
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

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

//leetcode submit region end(Prohibit modification and deletion)
