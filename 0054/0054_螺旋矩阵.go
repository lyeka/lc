package spiralMatrix

//leetcode submit region begin(Prohibit modification and deletion)
func spiralOrder(matrix [][]int) []int {
	ans := make([]int, 0)
	l, r, t, b := 0, len(matrix[0])-1, 0, len(matrix)-1
	for {
		for i := l; i <= r; i++ {
			ans = append(ans, matrix[t][i])
		}
		if t++; t > b {
			break
		}
		for i := t; i <= b; i++ {
			ans = append(ans, matrix[i][r])
		}
		if r--; r < l {
			break
		}
		for i := r; i >= l; i-- {
			ans = append(ans, matrix[b][i])
		}
		if b--; b < t {
			break
		}
		for i := b; i >= t; i-- {
			ans = append(ans, matrix[i][l])
		}
		if l++; l > r {
			break
		}
	}

	return ans

}

//leetcode submit region end(Prohibit modification and deletion)
