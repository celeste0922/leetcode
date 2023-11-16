package Array

//167. 两数之和II -输入有序数组 -Easy
//给你一个下标从1开始的整数数组numbers ，该数组已按非递减顺序排列  ，
//请你从数组中找出满足相加之和等于目标数target的两个数。
//如果设这两个数分别是numbers[index1]和numbers[index2] ，则1<=index1<index2<=numbers.length 。
//以长度为2的整数数组[index1, index2]的形式返回这两个整数的下标index1和index2。
//你可以假设每个输入只对应唯一的答案，而且你不可以重复使用相同的元素。
//你所设计的解决方案必须只使用常量级的额外空间。
//双指针

//	func TwoSum2(numbers []int, target int) []int {
//		n := len(numbers)
//		result := []int{numbers[0], numbers[0]}
//		for i := 0; i < n-1; i++ {
//			for j := i + 1; j < n; j++ {
//				flag := numbers[i] + numbers[j]
//				fmt.Println(flag)
//				if flag == target {
//					result[0], result[1] = i+1, j+1
//					return result
//				}
//				if flag > target {
//					n = j
//					break
//				}
//			}
//		}
//		return result
//	}
func TwoSum2(numbers []int, target int) []int {
	n := len(numbers)

	r := n - 1
	for l := 0; l <= r; {
		flag := numbers[l] + numbers[r]
		if flag == target {
			result := []int{l + 1, r + 1}
			return result
		}
		if flag > target {
			r--
		}
		if flag < target {
			l++
		}
	}
	return nil
}
