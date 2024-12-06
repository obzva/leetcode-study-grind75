/*
Big O
- N: the length of []int `prices`
- Time complexity: O(N)
- Space compleixty: O(1)
*/

func maxProfit(prices []int) int {
	minPrice := prices[0]
	profit := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
			continue
		}
		profit = max(profit, prices[i]-minPrice)
	}
	return profit
}
