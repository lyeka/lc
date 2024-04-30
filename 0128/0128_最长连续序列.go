package longestConsecutiveSequence

//leetcode submit region begin(Prohibit modification and deletion)
func longestConsecutive(nums []int) int {
	numSet := make(map[int]bool, len(nums))
	for _, num := range nums {
		numSet[num] = true
	}

	max := 0
	for num, _ := range numSet {
		if !numSet[num-1] {
			cnt := 1
			n := num
			for numSet[n+1] {
				cnt++
				n++
			}

			if cnt > max {
				max = cnt
			}
		}
	}

	return max
}

//leetcode submit region end(Prohibit modification and deletion)
