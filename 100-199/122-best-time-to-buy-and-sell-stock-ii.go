package main

func maxProfit(prices []int) int {
	var (
		totalProfit int
		buyPrice    = -1
	)
	if len(prices) == 1 {
		return 0
	}
	// 选择最初买入时间
	if prices[0] > prices[1] {
		buyPrice = prices[1]
	} else {
		buyPrice = prices[0]
	}

	for i := 1; i < len(prices)-1; i++ {
		if prices[i] > prices[i+1] {
			// Sell now & buy at next day
			totalProfit += prices[i] - buyPrice
			buyPrice = prices[i+1]
		}
	}
	// 最后一天可以考虑卖出
	if buyPrice < prices[len(prices)-1] {
		totalProfit += prices[len(prices)-1] - buyPrice
	}
	return totalProfit
}
