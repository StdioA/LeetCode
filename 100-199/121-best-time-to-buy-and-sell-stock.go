package main

func maxProfit(prices []int) int {
	var (
		minPrice  = prices[0]
		maxProfit int
	)
	for _, price := range prices[1:] {
		if price < minPrice {
			minPrice = price
		}
		if (price - minPrice) > maxProfit {
			maxProfit = price - minPrice
		}
	}
	return maxProfit
}
