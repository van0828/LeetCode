package main

import "fmt"

func main() {
	fmt.Println(maxProfit2([]int{1, 2, 3, 4, 5}))
}

//给定一个数组，它的第 i 个元素是一支给定股票第 i 天的价格。
//
//设计一个算法来计算你所能获取的最大利润。你可以尽可能地完成更多的交易（多次买卖一支股票）。
//
//注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。
//
//示例 1:
//
//输入: [7,1,5,3,6,4]
//输出: 7
//解释: 在第 2 天（股票价格 = 1）的时候买入，在第 3 天（股票价格 = 5）的时候卖出, 这笔交易所能获得利润 = 5-1 = 4 。
//     随后，在第 4 天（股票价格 = 3）的时候买入，在第 5 天（股票价格 = 6）的时候卖出, 这笔交易所能获得利润 = 6-3 = 3 。
//best-time-to-buy-and-sell-stock-ii
func maxProfit2(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	minPrice, profit := prices[0], 0
	for i := 1; i < len(prices); i++ {
		if i < len(prices)-1 && prices[i] > minPrice && prices[i] > prices[i+1] {
			profit += prices[i] - minPrice
			minPrice = prices[i+1]
			continue
		}
		if prices[i] < minPrice {
			minPrice = prices[i]
		}
		if i == len(prices)-1 && prices[i] > minPrice {
			profit += prices[i] - minPrice
		}
	}
	return profit
}
