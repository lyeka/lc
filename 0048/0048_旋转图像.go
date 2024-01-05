package rotateImage

//leetcode submit region begin(Prohibit modification and deletion)
func rotate(matrix [][]int) {
	rl := len(matrix) - 1
	for i := 0; i <= rl/2; i++ {
		for j := 0; j < len(matrix[0]); j++ {
			matrix[i][j], matrix[rl-i][j] = matrix[rl-i][j], matrix[i][j]
		}
	}

	for i := 1; i <= rl; i++ {
		for j := 0; j < i+1; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

}

//leetcode submit region end(Prohibit modification and deletion)
