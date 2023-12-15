package integerToRoman

//leetcode submit region begin(Prohibit modification and deletion)
func intToRoman(num int) string {
	list := []struct {
		value  int
		symbol string
	}{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}
	//res := ""
	//res := bytes.Buffer{}
	res := []byte{}
	for num > 0 {
		for _, n := range list {
			if n.value > num {
				continue
			}
			//res += n.symbol
			//res.WriteString(n.symbol)
			res = append(res, n.symbol...)
			num -= n.value
			break
		}
	}

	return string(res)

}

//leetcode submit region end(Prohibit modification and deletion)
