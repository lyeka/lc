package bestTimeToBuyAndSellStockIv

//leetcode submit region begin(Prohibit modification and deletion)
func maxProfit(k int, prices []int) int {

	buyK := make([]int, k+1)
	sellK := make([]int, k+1)
	for i := 1; i <= k; i++ { // 第一天无论第k手交易，都是最大利益都是-prices[0]
		buyK[i] = -prices[0]
	}

	for i := 0; i < len(prices); i++ {
		for j := 1; j <= k; j++ {
			buyK[j] = max(buyK[j], sellK[j-1]-prices[i])
			sellK[j] = max(sellK[j], buyK[j]+prices[i])
		}
	}

	return sellK[k]

}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

//leetcode submit region end(Prohibit modification and deletion)
