package Array

//766. 托普利茨矩阵 --Easy
//给你一个mxn的矩阵matrix 。如果这个矩阵是托普利茨矩阵，返回true ；否则，返回false 。
//如果矩阵上每一条由左上到右下的对角线上的元素都相同，那么这个矩阵是托普利茨矩阵 。
//==================================================================
//直接遍历全部元素除了左下右上

func IsToeplitzMatrix(matrix [][]int) bool {
	r := len(matrix)
	c := len(matrix[0])
	for i := 0; i < r-1; i++ {
		for j := 0; j < c-1; j++ {
			if matrix[i][j] != matrix[i+1][j+1] {
				return false
			}
		}
	}
	return true

}
