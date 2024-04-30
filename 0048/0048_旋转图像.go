package rotateImage

//leetcode submit region begin(Prohibit modification and deletion)
// 上下翻转+对角线翻转 = 旋转90°
// 上下翻转+左右翻转 = 旋转180°
func rotate(matrix [][]int) {
	rows := len(matrix) - 1
	// 上下翻转
	for r := 0; r <= rows/2; r++ {
		for c := 0; c < len(matrix[0]); c++ {
			matrix[r][c], matrix[rows-r][c] = matrix[rows-r][c], matrix[r][c]
		}
	}

	// 对角线翻转
	for r := 1; r <= rows; r++ {
		for c := 0; c < r+1; c++ {
			matrix[r][c], matrix[c][r] = matrix[c][r], matrix[r][c]
		}
	}

}

//leetcode submit region end(Prohibit modification and deletion)
