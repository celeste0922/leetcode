package Array

import "sort"

//628. 三个数的最大乘积  --Easy
//给你一个整型数组nums ，在数组中找出由三个数组成的最大乘积，并输出这个乘积。
//===============================================================
//1.求最大的三个整数，最小的两个负数*最大的正数，乘积最大的即为结果
//2.求最大的三个正数和最小的两个负数即可。

func MaximumProduct(nums []int) int {
	sort.Ints(nums)
	n := len(nums)
	return max(nums[0]*nums[1]*nums[n-1], nums[n-1]*nums[n-2]*nums[n-3])

}
