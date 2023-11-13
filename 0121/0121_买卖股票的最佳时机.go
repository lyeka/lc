package bestTimeToBuyAndSellStock

//leetcode submit region begin(Prohibit modification and deletion)
func maxProfit(prices []int) int {
	minPrice := prices[0]
	maxProfit := 0
	for _, price := range prices { // 每一天都和该天之前的最低价格比较
		maxProfit = max(maxProfit, price-minPrice)
		minPrice = min(minPrice, price)

	}
	return maxProfit

}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

//leetcode submit region end(Prohibit modification and deletion)
