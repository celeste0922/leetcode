package Array

//54. 螺旋矩阵
//给你一个 m 行 n 列的矩阵 matrix ，请按照 顺时针螺旋顺序 ，返回矩阵中的所有元素。
//======================================================================
//先找到四个角的边界，然后按方向拆分为四个有序执行的遍历步骤
//四个边界逐渐收缩

func SpiralOrder(matrix [][]int) []int {
	y := len(matrix)
	x := len(matrix[0])
	index := 0
	result := make([]int, x*y)
	left, right, top, bottom := 0, x-1, 0, y-1
	columns, rows := left, top
	for left <= right && top <= bottom {
		for columns := left; columns <= right; columns++ {
			result[index] = matrix[top][columns]
			index++
		}
		for rows := top + 1; rows <= bottom; rows++ {
			result[index] = matrix[rows][right]
			index++
		}
		if left < right && top < bottom { //当left==right||top==bottom时，只需要遍历一或两格，前面已经执行完，后面无需执行
			for columns = right - 1; columns > left; columns-- {
				result[index] = matrix[bottom][columns]
				index++
			}
			for rows = bottom; rows > top; rows-- {
				result[index] = matrix[rows][left]
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
