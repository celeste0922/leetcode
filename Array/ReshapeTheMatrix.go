package Array

//566. 重塑矩阵   --Easy
//在MATLAB中，有一个非常有用的函数reshape ，它可以将一个m*n矩阵重塑为另一个大小不同（r*c）的新矩阵，但保留其原始数据。
//给你一个由二维数组mat表示的m*n矩阵，以及两个正整数r和c ，分别表示想要的重构的矩阵的行数和列数。
//重构后的矩阵需要将原始矩阵的所有元素以相同的行遍历顺序填充。
//如果具有给定参数的reshape操作是可行且合理的，则输出新的重塑矩阵；否则，输出原始矩阵。
//===========================================================================
//数学！数学!数学！还是他妈的数学！

func MatrixReshape(mat [][]int, r int, c int) [][]int {
	n, m := len(mat), len(mat[0])
	if n*m != r*c {
		return mat
	}
	ans := make([][]int, r)
	for key := range ans {
		ans[key] = make([]int, c)
	}
	for i := 0; i < m*n; i++ {
		ans[i/c][i%c] = mat[i/m][i%m]
	}
	return ans
}
