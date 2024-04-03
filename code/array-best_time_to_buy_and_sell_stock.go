package main

func maxProfit(prices []int) int {
	maxProfit := 0
	buyDay := 0
	for today, price := range prices {
		if price < prices[buyDay] {
			buyDay = today
		}
		profit := prices[today] - prices[buyDay]
		maxProfit = max(maxProfit, profit)
	}

	return maxProfit
}
