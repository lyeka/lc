package perfectSquares

//leetcode submit region begin(Prohibit modification and deletion)
func numSquares(n int) int {
	f := make([]int, n+1)
	for i := 1; i <= n; i++ {
		minn := n + 1
		for j := 1; j*j <= i; j++ {
			minn = min(minn, f[i-j*j])
		}
		f[i] = minn + 1
	}

	return f[n]

}

func min(x, y int) int {
	if x < y {
		return x
	}

	return y
}

//leetcode submit region end(Prohibit modification and deletion)
