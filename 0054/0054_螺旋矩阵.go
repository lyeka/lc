package spiralMatrix

//leetcode submit region begin(Prohibit modification and deletion)
func spiralOrder(matrix [][]int) []int {
	l, r, t, b := 0, len(matrix[0])-1, 0, len(matrix)-1
	res := []int{}
	for {
		for i := l; i <= r; i++ {
			res = append(res, matrix[t][i])
		}
		if t++; t > b {
			break
		}

		for i := t; i <= b; i++ {
			res = append(res, matrix[i][r])
		}
		if r--; r < l {
			break
		}

		for i := r; i >= l; i-- {
			res = append(res, matrix[b][i])
		}
		if b--; b < t {
			break
		}

		for i := b; i <= t; i-- {
			res = append(res, matrix[i][l])
		}
		if l++; l > r {
			break
		}
	}

	return res

}

//leetcode submit region end(Prohibit modification and deletion)
