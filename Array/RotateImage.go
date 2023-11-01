package Array

//48. 旋转图像
//给定一个 n × n 的二维矩阵 matrix 表示一个图像。请你将图像顺时针旋转 90 度。
//你必须在 原地 旋转图像，这意味着你需要直接修改输入的二维矩阵。请不要 使用另一个矩阵来旋转图像。
//=============================================================================
//作为例子，分析将图像旋转 90 度之后，这些数字出现在什么位置。
//对于矩阵中的第一行而言，在旋转后，它出现在倒数第一列的位置。
//并且，第一行的第 xxx 个元素在旋转后恰好是倒数第一列的第 xxx 个元素。
//对于矩阵中的第二行而言，在旋转后，它出现在倒数第二列的位置。
//====================================================
//将矩阵分为四部分，每部分旋转寻找每部分旋转后的落点规律
//(i,j) -> (j, length - 1 - i) -> (length - 1 - i, length - 1 - j) -> (length - 1 - j, i) -> (i,j)
//矩阵杀我

func Rotate(matrix [][]int) [][]int {
	n := len(matrix)
	for i := 0; i < n/2; i++ {
		for j := 0; j < (n+1)/2; j++ {
			matrix[i][j], matrix[n-j-1][i], matrix[n-i-1][n-j-1], matrix[j][n-i-1] =
				matrix[n-j-1][i], matrix[n-i-1][n-j-1], matrix[j][n-i-1], matrix[i][j]
		}
	}
	return matrix
}
