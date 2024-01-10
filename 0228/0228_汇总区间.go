package summaryRanges

import (
	"fmt"
)

//leetcode submit region begin(Prohibit modification and deletion)
func summaryRanges(nums []int) []string {
	res := []string{}
	index := 0
	for index < len(nums) {
		x := 1
		for index+x < len(nums) && nums[index+x] == nums[index]+1*x {
			x++
		}
		if x > 1 {
			res = append(res, fmt.Sprintf("%d->%d", nums[index], nums[index+x-1]))
		} else {
			res = append(res, fmt.Sprintf("%d", nums[index]))
		}

		index += x

	}

	return res

}

//leetcode submit region end(Prohibit modification and deletion)
