package mergeSortedArray

//leetcode submit region begin(Prohibit modification and deletion)
func merge(nums1 []int, m int, nums2 []int, n int) {
	index := m + n - 1
	i, j := m-1, n-1
	for i >= 0 && j >= 0 {
		if nums1[i] > nums2[j] {
			nums1[index] = nums1[i]
			i--
		} else {
			nums1[index] = nums2[j]
			j--
		}
		index--
	}
	for i >= 0 {
		nums1[index] = nums1[i]
		i--
		index--

	}

	for j >= 0 {
		nums1[index] = nums2[j]
		j--
		index--
	}
}

//leetcode submit region end(Prohibit modification and deletion)
