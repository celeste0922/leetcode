package Array

import "fmt"

// 122. 买卖股票的最佳时机II
// 给你一个整数数组prices，其中prices[i]表示某支股票第i天的价格。
// 在每一天，你可以决定是否购买和/或出售股票。你在任何时候最多只能持有一股股票。你也可以先购买，然后在同一天出售。
// 返回你能获得的最大利润。
// =================================
//动态规划=>状态：持有\\未持有

func MaxProfit2(prices []int) int {
	dp := make([][]int, 2)
	for key := range dp {
		dp[key] = make([]int, 2)
	}
	dp[0][0], dp[1][0] = -prices[0], -prices[0]
	dp[0][1], dp[1][1] = 0, 0
	fmt.Println(dp)
	for i := 1; i < len(prices); i++ {
		dp[0][0], dp[0][1] = dp[1][0], dp[1][1]
		dp[1][0] = max(dp[0][1]-prices[i], dp[0][0])
		dp[1][1] = max(dp[0][1], prices[i]+dp[0][0])
		fmt.Println(dp)
	}
	return dp[1][1]
}

//====================贪心=========================
//从第二天开始，如果当天股价大于前一天股价，则在前一天买入，当天卖出，即可获得利润。
//如果当天股价小于前一天股价，则不买入，不卖出。
//所有上涨交易日都做买卖，所有下跌交易日都不做买卖，最终获得的利润是最大的。

//func maxProfit(prices []int) (ans int) {
//	for i, v := range prices[1:] {
//		t := v - prices[i]
//		if t > 0 {
//			ans += t
//		}
//	}
//	return
//}
