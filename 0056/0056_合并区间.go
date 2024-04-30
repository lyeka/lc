package mergeIntervals

import (
	"sort"
)

//leetcode submit region begin(Prohibit modification and deletion)
func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	ans := [][]int{intervals[0]}

	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] <= ans[len(ans)-1][1] {
			if intervals[i][1] > ans[len(ans)-1][1] {
				ans[len(ans)-1][1] = intervals[i][1]
			}
		} else {
			ans = append(ans, intervals[i])
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
