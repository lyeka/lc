package mergeSortedArray

//leetcode submit region begin(Prohibit modification and deletion)
func merge(nums1 []int, m int, nums2 []int, n int) {
	sorted := make([]int, 0, m+n)
	i, j := 0, 0
	for {
		if i == m {
			sorted = append(sorted, nums2[j:]...)
			break
		}
		if j == n {
			sorted = append(sorted, nums1[i:]...)
			break
		}
		if nums1[i] <= nums2[j] {
			sorted = append(sorted, nums1[i])
			i++
		} else {
			sorted = append(sorted, nums2[j])
			j++
		}
	}
	copy(nums1, sorted)
}

//leetcode submit region end(Prohibit modification and deletion)
