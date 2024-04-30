package subarraySumEqualsK

//leetcode submit region begin(Prohibit modification and deletion)
func subarraySum(nums []int, k int) int {
	cnt := 0
	pre := 0
	m := map[int]int{}
	//m[0] = 1 可以替代 pre - k == 0  这个判断
	for i := 0; i < len(nums); i++ {
		pre += nums[i]
		if _, ok := m[pre-k]; ok {
			cnt += m[pre-k]
		}
		if pre-k == 0 {
			cnt++
		}

		m[pre]++

	}

	return cnt
}

//leetcode submit region end(Prohibit modification and deletion)
