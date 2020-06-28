package main

import "fmt"

func main() {
	fmt.Println(maxProfit([]int{7,6,4,3,1}))
}

//给定一个数组，它的第 i 个元素是一支给定股票第 i 天的价格。
//
//如果你最多只允许完成一笔交易（即买入和卖出一支股票），设计一个算法来计算你所能获取的最大利润。
//
//注意你不能在买入股票前卖出股票。
//
//示例 1:
//
//输入: [7,1,5,3,6,4]
//输出: 5
//解释: 在第 2 天（股票价格 = 1）的时候买入，在第 5 天（股票价格 = 6）的时候卖出，最大利润 = 6-1 = 5 。
//     注意利润不能是 7-1 = 6, 因为卖出价格需要大于买入价格。
// best-time-to-buy-and-sell-stock
func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	profit, minPrice := 0, prices[0]
	for i := 1; i < len(prices); i++ {
		if prices[i] - minPrice > profit {
			profit = prices[i] - minPrice
		}
		if prices[i] < minPrice {
			minPrice = prices[i]
		}
	}
	return profit
}
