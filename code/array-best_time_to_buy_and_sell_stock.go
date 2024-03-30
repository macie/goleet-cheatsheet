package main

func maxProfit(prices []int) int {
	minPrice := 10000 + 1

	profit := 0
	for _, price := range prices {
		if price < minPrice {
			minPrice = price
		} else if price-minPrice > profit {
			profit = price - minPrice
		}
	}

	return profit
}
