package Array

//63. 不同路径II
//一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为 “Start” ）。
//机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为 “Finish”）。
//现在考虑网格中有障碍物。那么从左上角到右下角将会有多少条不同的路径？
//网格中的障碍物和空位置分别用 1 和 0 来表示。
//==============================================================================
//还是动态规划，使用原数组。
//读取到当前节点时，判断当前节点是否==0，不==0则是障碍物，将当前节点值置0（意味着后面计算时此路不通：路径数为0）
//边界条件是真他吗多啊
//起点和终点为1时无路径
//只有一列或一行时路径数为1
//边缘有障碍时障碍点及后面的边经过路径都为0
//====================================
//官方题解使用新的数组，数组中全是0，逻辑更为简单

func UniquePathsWithObstacles(obstacleGrid [][]int) int {
	x := len(obstacleGrid[0])
	y := len(obstacleGrid)
	p := 1
	q := 1
	if obstacleGrid[y-1][x-1] == 1 {
		return 0
	}
	if obstacleGrid[0][0] == 1 {
		return 0
	}
	if x == 1 || y == 1 {
		return 1
	}
	for i := 0; i < x; i++ {
		if obstacleGrid[0][i] == 1 { //产生障碍物，则后面赋值0
			p = 0
		}
		obstacleGrid[0][i] = p
	}
	for i := 1; i < y; i++ {
		if obstacleGrid[i][0] == 1 {
			q = 0
		}
		obstacleGrid[i][0] = q
	}
	for c := 1; c < y; c++ {
		for r := 1; r < x; r++ {
			if obstacleGrid[c][r] == 1 { //当前节点是障碍物，则通过路径是0
				obstacleGrid[c][r] = 0
				continue
			}
			obstacleGrid[c][r] = obstacleGrid[c-1][r] + obstacleGrid[c][r-1]
		}
	}
	return obstacleGrid[y-1][x-1]
}
