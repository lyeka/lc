package setMatrixZeroes

//leetcode submit region begin(Prohibit modification and deletion)
func setZeroes(matrix [][]int) {
	row0, col0 := false, false
	m := len(matrix)
	n := len(matrix[0])

	// 判断第一行是否有0
	for i := 0; i < n; i++ {
		if matrix[0][i] == 0 {
			row0 = true
			break
		}
	}
	// 判断第一列是否有0
	for i := 0; i < m; i++ {
		if matrix[i][0] == 0 {
			col0 = true
			break
		}
	}

	// 判断其余行列是否有0，使用原矩阵的第一行第一列做标记
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][j] == 0 {
				matrix[0][j] = 0
				matrix[i][0] = 0
			}
		}
	}

	// 使用其余行列的标记置0
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[0][j] == 0 || matrix[i][0] == 0 {
				matrix[i][j] = 0
			}
		}
	}

	// 第一行置0
	if row0 {
		for i := 0; i < n; i++ {
			matrix[0][i] = 0
		}
	}

	// 第一列置0
	if col0 {
		for i := 0; i < m; i++ {
			matrix[i][0] = 0
		}
	}

}

//leetcode submit region end(Prohibit modification and deletion)

// 常规做法，空间复杂度：O(m+n)
func setZeroesV2(matrix [][]int) {
	row := make([]int, len(matrix))
	col := make([]int, len(matrix[0]))
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == 0 {
				row[i] = 1
				col[j] = 1
			}
		}
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if row[i] == 1 || col[j] == 1 {
				matrix[i][j] = 0
			}
		}
	}

	return

}
