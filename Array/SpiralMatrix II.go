package Array

//59. 螺旋矩阵II
//给你一个正整数 n ，生成一个包含 1 到 n2 所有元素，且元素按顺时针顺序螺旋排列的 n x n 正方形矩阵 matrix 。
//直接引用螺旋矩阵 哈哈哈

func GenerateMatrix(n int) [][]int {
	result := make([][]int, n)
	for i := 0; i < n; i++ {
		result[i] = make([]int, n)
	}
	num := 1
	index := 0
	left, right, top, bottom := 0, n-1, 0, n-1
	columns, rows := left, top
	for left <= right && top <= bottom {
		for columns = left; columns <= right; columns++ {
			result[top][columns] = num
			num++
			index++
		}
		for rows = top + 1; rows <= bottom; rows++ {
			result[rows][right] = num
			num++
			index++
		}
		if left < right && top < bottom { //当left==right||top==bottom时，只需要遍历一或两格，前面已经执行完，后面无需执行
			for columns = right - 1; columns > left; columns-- {
				result[bottom][columns] = num
				num++
				index++
			}
			for rows = bottom; rows > top; rows-- {
				result[rows][left] = num
				num++
				index++
			}
		}
		left++
		right--
		top++
		bottom--
	}
	return result
}
