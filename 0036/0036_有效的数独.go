package validSudoku

//leetcode submit region begin(Prohibit modification and deletion)
func isValidSudoku(board [][]byte) bool {
	var rows, columns = [9][9]int{}, [9][9]int{}
	var subbox = [3][3][9]int{}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == '.' {
				continue
			}
			index := board[i][j] - '1'
			rows[i][index]++
			columns[j][index]++
			subbox[i/3][j/3][index]++
			if rows[i][index] > 1 || columns[j][index] > 1 || subbox[i/3][j/3][index] > 1 {
				return false
			}
		}
	}

	return true

}

//leetcode submit region end(Prohibit modification and deletion)
