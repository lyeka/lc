package insertInterval

//leetcode submit region begin(Prohibit modification and deletion)
func insert(intervals [][]int, newInterval []int) [][]int {
	res := [][]int{}
	merge := false
	left, right := newInterval[0], newInterval[1]

	for _, interval := range intervals {
		if interval[0] > right {
			if !merge {
				res = append(res, []int{left, right})
				merge = true
			}
			res = append(res, interval)
		} else if interval[1] < left {
			res = append(res, interval)
		} else {
			left = min(left, interval[0])
			right = max(right, interval[1])
		}
	}

	if !merge {
		res = append(res, []int{left, right})
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
