package threeSum

import (
	"sort"
)

//leetcode submit region begin(Prohibit modification and deletion)
func threeSum(nums []int) [][]int {
	sort.Ints(nums)

	ans := make([][]int, 0)
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			return ans
		}

		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		l, r := i+1, len(nums)-1
		for l < r {
			n := nums[i] + nums[l] + nums[r]
			if n == 0 {
				ans = append(ans, []int{nums[i], nums[l], nums[r]})
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
			if n < 0 {
				l++
			}
			if n > 0 {
				r--
			}
		}
	}

	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
