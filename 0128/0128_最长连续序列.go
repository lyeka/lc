package longestConsecutiveSequence

//leetcode submit region begin(Prohibit modification and deletion)
func longestConsecutive(nums []int) int {
	numSet := map[int]bool{}
	for _, num := range nums {
		numSet[num] = true
	}

	longest := 0
	for num, _ := range numSet {
		if !numSet[num-1] {
			cnt := 1
			tn := num
			for numSet[tn+1] {
				cnt++
				tn++
			}
			if cnt > longest {
				longest = cnt
			}
		}
	}

	return longest

}

//leetcode submit region end(Prohibit modification and deletion)
