package Array

import "fmt"

//121. 买卖股票的最佳时机
//给定一个数组 prices ，它的第i个元素prices[i]表示一支给定股票第i天的价格。
//你只能选择某一天买入这只股票，并选择在 未来的某一个不同的日子卖出该股票。设计一个算法来计算你所能获取的最大利润。
//返回你可以从这笔交易中获取的最大利润。如果你不能获取任何利润，返回0。
//直接记录最小成本和最大利润往后遍历
//用动态规划有点折磨人

func MaxProfit(prices []int) int {
	dp := make([][]int, 2)
	for key := range dp {
		dp[key] = make([]int, 2)
	}
	dp[0][0], dp[1][0] = -prices[0], -prices[0]
	dp[0][1], dp[1][1] = 0, 0
	fmt.Println(dp)
	for i := 1; i < len(prices); i++ {
		dp[0][0], dp[0][1] = dp[1][0], dp[1][1]
		dp[1][0], dp[1][1] = max(-prices[i], dp[0][0]), max(dp[0][1], prices[i]+dp[0][0])
		fmt.Println(dp)
	}
	return dp[1][1]
}

//从左到右枚举卖出价格prices[i]，那么要想获得最大利润，我们需要知道第i天之前，股票价格的最小值是什么，
//也就是从 prices[0]到 prices[i−1]的最小值，把它作为买入价格，这可以用一个变量 minPrice维护。
//请注意，minPrice维护的是 prices[i]左侧元素的最小值。
//由于只能买卖一次，所以在遍历中，维护 prices[i]−minPrice的最大值，就是答案。
//=====================================================================
//func maxProfit(prices []int) (ans int) {
//	minPrice := prices[0]
//	for _, p := range prices {
//		ans = max(ans, p-minPrice)
//		minPrice = min(minPrice, p)
//	}
//	return
//}
