package Array

//152. 乘积最大子数组
//给你一个整数数组 nums ，请你找出数组中乘积最大的非空连续子数组（该子数组中至少包含一个数字），并返回该子数组所对应的乘积。
//测试用例的答案是一个 32-位 整数。
//子数组 是数组的连续子序列。
//==============================
//动态规划
//考虑当前最大值的可能性有：当前元素的值，当前元素的值*之前的最大值，当前元素的值*之前的最小值（负数相乘）
//维护三个值，其中两个是状态=>历史最大值，历史最小值；然后是结果res=max(res, maxNum)

func MaxProduct(nums []int) int {
	n := len(nums)
	minNum, maxNum, res := nums[0], nums[0], nums[0]
	if n == 1 {
		return nums[0]
	}
	for i := 1; i < n; i++ {
		minNum, maxNum = min(min(minNum*nums[i], nums[i]), maxNum*nums[i]), max(max(maxNum*nums[i], nums[i]), minNum*nums[i])
		res = max(res, maxNum)
	}
	return res
}
