package Array

//双指针
//如果移动数字较大的那个指针，那么前者「两个指针指向的数字中较小值」不会增加
//后者「指针之间的距离」会减小，那么这个乘积会减小。因此，移动数字较大的那个指针是不合理的。
//于是，移动 数字较小的那个指针。

func MaxArea(height []int) int {
	i := 0
	j := len(height) - 1
	area := 0
	for i != j {
		area = max(area, (j-i)*min(height[i], height[j]))
		if height[i] > height[j] {
			j--
		} else {
			i++
		}
	}
	return area
}
func max(a int, b int) int {
	if a > b {
		return a
	}
	return b

}
func min(a int, b int) int {
	if a < b {
		return a
	}
	return b

}
