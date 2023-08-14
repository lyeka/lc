package climbingStairs

//leetcode submit region begin(Prohibit modification and deletion)
func climbStairs(n int) int {
	p, q, r := 0, 0, 1
	for i := 0; i < n; i++ {
		p = q
		q = r
		r = p + q

	}

	return r

}

//leetcode submit region end(Prohibit modification and deletion)

// 递归版本，会超时
//func climbStairs(n int) int {
//	if n <= 2 {
//		return n
//	}
//	return climbStairs(n-1) + climbStairs(n-2)
//
//}
