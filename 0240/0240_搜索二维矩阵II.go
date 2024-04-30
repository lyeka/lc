package searchA2dMatrixIi

//leetcode submit region begin(Prohibit modification and deletion)
func searchMatrix(matrix [][]int, target int) bool {
	var left, right, mid int
	for i := len(matrix) - 1; i >= 0; i-- {
		if matrix[i][0] > target {
			continue
		} else {
			left, right = 0, len(matrix[i])-1
			for left <= right {
				mid = (left + right) / 2
				if matrix[i][mid] == target {
					return true
				}
				if matrix[i][mid] < target {
					left = mid + 1
				} else {
					right = mid - 1
				}
			}
		}
	}
	return false

}

//leetcode submit region end(Prohibit modification and deletion)

func searchMatrixV2(matrix [][]int, target int) bool {
	for i := len(matrix) - 1; i >= 0; i-- {
		if matrix[i][0] > target {
			continue
		} else {
			for j := 0; j < len(matrix[i]); j++ {
				if matrix[i][j] == target {
					return true
				}
			}
		}
	}

	return false

}
