package pascalsTriangle

//leetcode submit region begin(Prohibit modification and deletion)
func generate(numRows int) [][]int {
	ans := make([][]int, numRows)
	ans[0] = []int{1}

	for i := 1; i < numRows; i++ {
		tmp := make([]int, i+1)
		tmp[0] = 1
		tmp[i] = 1
		for j := 1; j < i; j++ {
			tmp[j] = ans[i-1][j] + ans[i-1][j-1]
		}
		ans[i] = tmp
	}

	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
