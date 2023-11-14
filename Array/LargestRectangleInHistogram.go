package Array

//84. 柱状图中最大的矩形
//给定 n 个非负整数，用来表示柱状图中各个柱子的高度。每个柱子彼此相邻，且宽度为 1 。
//求在该柱状图中，能够勾勒出来的矩形的最大面积。
//=====================================================================
//太哈人了单调栈
//考虑记录一下左边界或者右边界，这样就不用每次寻找了，
//我们就可以采用单调栈，为一个递增的单调栈，
//当当前访问的柱子比栈顶元素的柱子大的时候，就说明我们找到了右边界（就是当前柱子下标），
//就可以求解栈顶柱子的矩形面积了，左边界呢？因为是递增的，
//使用左边界就是栈顶元素的后一个元素（也就是比他先一步入栈的那个元素），高就是栈顶柱子的高度， 面积就不得而知了
//当单调栈弹出元素时，弹出的这个元素可以计算面积（其右边界另其弹出，左边界为左边相邻第一个元素）

func LargestRectangleArea(heights []int) int {
	//单调栈（单调递增）
	stack := make([]int, 0)
	stack = append(stack, -1) //stack的哨兵，方便确定左边界
	//添加一个哨兵，减少代码量，这个哨兵可以令最后一个元素弹出，此时计算最后一个元素的面积
	heights = append(heights, 0)
	ln := len(heights)
	res := 0 //结果
	for i := 0; i < ln; i++ {
		for len(stack) > 1 && heights[i] < heights[stack[len(stack)-1]] {
			res = max(res, heights[stack[len(stack)-1]]*(i-stack[len(stack)-2]-1))
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	return res
}

//func max(a int, b int) int {
//	if a > b {
//		return a
//	}
//	return b
//
//}
