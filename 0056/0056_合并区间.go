package mergeIntervals

import (
	"sort"
)

//leetcode submit region begin(Prohibit modification and deletion)
func merge(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	merge := [][]int{intervals[0]}
	mIndex := 0

	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] > merge[mIndex][1] {
			merge = append(merge, intervals[i])
			mIndex++
		} else {
			merge[mIndex][1] = max(intervals[i][1], merge[mIndex][1])
		}
	}

	return merge

}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

//leetcode submit region end(Prohibit modification and deletion)
