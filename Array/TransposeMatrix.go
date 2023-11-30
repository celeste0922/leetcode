package Array

//867. 转置矩阵
//给你一个二维整数数组 matrix， 返回 matrix 的 转置矩阵 。
//矩阵的 转置 是指将矩阵的主对角线翻转，交换矩阵的行索引与列索引。
//===================================================

func Transpose(matrix [][]int) [][]int {
	m, n := len(matrix), len(matrix[0])
	if m == n {
		// 原地转置
		for i := 0; i < n; i++ {
			for j := i + 1; j < m; j++ {
				matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
			}
		}
		return matrix
	} else {
		// 构造新数组转置
		ans := make([][]int, 0, n)
		for i := 0; i < n; i++ {
			ans = append(ans, make([]int, 0, m))
			for j := 0; j < m; j++ {
				ans[i] = append(ans[i], matrix[j][i])
			}
		}
		return ans
	}

}
