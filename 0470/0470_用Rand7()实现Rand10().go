package implementRand10UsingRand7

//leetcode submit region begin(Prohibit modification and deletion)
func rand10() int {
	tmp := 0
	for i := 1; i <= 10; i++ {
		tmp += rand7()
	}

	return tmp%10 + 1

}

//leetcode submit region end(Prohibit modification and deletion)
