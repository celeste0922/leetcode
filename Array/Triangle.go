package Array

import "sort"

//120. 三角形最小路径和
//给定一个三角形triangle，找出自顶向下的最小路径和。
//每一步只能移动到下一行中相邻的结点上。相邻的结点在这里指的是下标与上一层结点下标相同或者等于上一层结点下标+1的两个结点。
//也就是说，如果正位于当前行的下标i，那么下一步可以移动到下一行的下标i或i+1 。
//=================================================================
//动态规划

func MinimumTotal(triangle [][]int) int {
	n := len(triangle)
	for i := 1; i < n; i++ {
		l := len(triangle[i])
		triangle[i][0] += triangle[i-1][0]
		triangle[i][l-1] += triangle[i-1][l-2]
		for j := 1; j < l-1; j++ {
			triangle[i][j] += min(triangle[i-1][j], triangle[i-1][j-1])
		}
	}
	sort.Ints(triangle[n-1])
	return triangle[n-1][0]
}

//func min(a int, b int) int {
//	if a < b {
//		return a
//	}
//	return b
//
//}
