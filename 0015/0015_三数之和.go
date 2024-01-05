package threeSum

import (
	"sort"
)

//leetcode submit region begin(Prohibit modification and deletion)
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	res := make([][]int, 0)

	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			continue
		}
		if i > 0 && nums[i-1] == nums[i] {
			continue
		}
		l, r := i+1, len(nums)-1
		for l < r {
			t := nums[i] + nums[l] + nums[r]
			if t == 0 {
				res = append(res, []int{nums[i], nums[l], nums[r]})
				for l < r && nums[l] == nums[l+1] {
					l++
				}
				for l < r && nums[r] == nums[r-1] {
					r--
				}
				l++
				r--
				continue
			}
			if t < 0 {
				l++
			}
			if t > 0 {
				r--
			}
		}
	}

	return res
}

//leetcode submit region end(Prohibit modification and deletion)
