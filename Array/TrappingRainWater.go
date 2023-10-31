package Array

//42.接雨水
//给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。
//========================================================================
//我的思路=>从第一个开始，找最高，找到后原路径计算每一列水的容积，依次往下找（暴力）
//双指针
//=====================================================================
//在暴力解法中，我们可以看到只要记录左边柱子的最高高度 和 右边柱子的最高高度，就可以计算当前位置的雨水面积，这就是通过列来计算。
//当前列雨水面积：min(左边柱子的最高高度，记录右边柱子的最高高度) - 当前柱子高度。
//为了得到两边的最高高度，使用了双指针来遍历，每到一个柱子都向两边遍历一遍，这其实是有重复计算的。
//我们把每一个位置的左边最高高度记录在一个数组上（maxLeft），右边最高高度记录在一个数组上（maxRight），这样就避免了重复计算。
//当前位置，左边的最高高度是前一个位置的左边最高高度和本高度的最大值。
//即从左向右遍历：maxLeft[i] = max(height[i], maxLeft[i - 1]);
//从右向左遍历：maxRight[i] = max(height[i], maxRight[i + 1]);

func Trap(height []int) int {
	n := len(height)
	size := 0
	lheight := make([]int, n)
	rheight := make([]int, n)
	lheight[0] = height[0]
	rheight[n-1] = height[n-1]
	for index := 1; index < n; index++ {
		lheight[index] = max(lheight[index-1], height[index])
	}
	for index := n - 2; index > -1; index-- {
		rheight[index] = max(rheight[index+1], height[index])
	}
	for i := 1; i < n-1; i++ {
		size += min(rheight[i], lheight[i]) - height[i]
	}
	return size
}

//func max(a int, b int) int {
//	if a > b {
//		return a
//	}
//	return b
//
//}
//func min(a int, b int) int {
//	if a < b {
//		return a
//	}
//	return b
//}
