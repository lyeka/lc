package slidingWindowMaximum

//leetcode submit region begin(Prohibit modification and deletion)

func maxSlidingWindow(nums []int, k int) []int {
	q := []int{}
	push := func(i int) {
		for len(q) > 0 && nums[i] >= nums[q[len(q)-1]] {
			q = q[:len(q)-1]
		}
		q = append(q, i)
	}
	for i := 0; i < k; i++ {
		push(i)
	}
	ans := make([]int, 0, len(nums)-k)
	ans = append(ans, nums[q[0]])
	for i := k; i < len(nums); i++ {
		push(i)
		for q[0] <= i-k {
			q = q[1:]
		}
		ans = append(ans, nums[q[0]])
	}

	return ans

}

//leetcode submit region end(Prohibit modification and deletion)
