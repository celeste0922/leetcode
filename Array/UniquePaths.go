package Array

// 62. 不同路径
// 一个机器人位于一个 m*n 网格的左上角 （起始点在下图中标记为“Start”）。
// 机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为“Finish”）。
// 问总共有多少条不同的路径？
// ===========================
// 快使用动态规划
//动态规划
//机器人从(0 , 0) 位置出发，到(m - 1, n - 1)终点。
//按照动规五部曲来分析：
//确定dp数组（dp table）以及下标的含义
//dp[i][j] ：表示从（0 ，0）出发，到(i, j) 有dp[i][j]条不同的路径。
//确定递推公式
//想要求dp[i][j]，只能有两个方向来推导出来，即dp[i - 1][j] 和 dp[i][j - 1]。
//此时在回顾一下 dp[i - 1][j] 表示啥，是从(0, 0)的位置到(i - 1, j)有几条路径，dp[i][j - 1]同理。
//那么很自然，dp[i][j] = dp[i - 1][j] + dp[i][j - 1]，因为dp[i][j]只有这两个方向过来。
//dp数组的初始化
//如何初始化呢，首先dp[i][0]一定都是1，因为从(0, 0)的位置到(i, 0)的路径只有一条，那么dp[0][j]也同理。
//所以初始化代码为：
//for (int i = 0; i < m; i++) dp[i][0] = 1;
//for (int j = 0; j < n; j++) dp[0][j] = 1;
//确定遍历顺序
//这里要看一下递推公式dp[i][j] = dp[i - 1][j] + dp[i][j - 1]，dp[i][j]都是从其上方和左方推导而来，那么从左到右一层一层遍历就可以了。
//这样就可以保证推导dp[i][j]的时候，dp[i - 1][j] 和 dp[i][j - 1]一定是有数值的。

func UniquePaths(m int, n int) int {
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		dp[i][0] = 1
	}
	for i := 0; i < n; i++ {
		dp[0][i] = 1
	}

	for y := 1; y < m; y++ {
		for x := 1; x < n; x++ {
			dp[y][x] = dp[y-1][x] + dp[y][x-1]
		}
	}
	return dp[m-1][n-1]
}

//ds已然寄咯
//func UniquePaths(m int, n int) int {
//	nums = 0
//	path := make([][]bool, m)
//	for i := 0; i < m; i++ {
//		path[i] = make([]bool, n)
//	}
//	walk(path, 0, 0)
//	return nums
//}
//
//var nums = 0
//
//func walk(n [][]bool, x int, y int) { //我的深度搜索=>果不其然超时力
//	if x == len(n[0])-1 && y == len(n)-1 {
//		nums++
//		return
//	}
//	if x+1 < len(n[0]) {
//		walk(n, x+1, y)
//	}
//	if y+1 < len(n) {
//		walk(n, x, y+1)
//	}
//}
