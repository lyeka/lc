package zigzagConversion

//leetcode submit region begin(Prohibit modification and deletion)
func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	res := make([][]byte, numRows)

	n := -1
	flag := 1
	for i := 0; i < len(s); i++ {
		n += flag
		if n == numRows-1 {
			flag = -1
		}
		if n == 0 {
			flag = 1
		}
		res[n] = append(res[n], s[i])
	}

	resStr := ""
	for i := 0; i < numRows; i++ {
		resStr += string(res[i])
	}

	return resStr

}

//leetcode submit region end(Prohibit modification and deletion)
