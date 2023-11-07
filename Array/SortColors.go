package Array

import "fmt"

// 75. 颜色分类
// 给定一个包含红色、白色和蓝色、共 n 个元素的数组 nums ，
// 原地对它们进行排序，使得相同颜色的元素相邻，并按照红色、白色、蓝色顺序排列。
// 我们使用整数 0、 1 和 2 分别表示红色、白色和蓝色。
// 必须在不使用库内置的 sort 函数的情况下解决这个问题。
// ===============================================
// 经典双指针
//整体思路是将所有的0放到左边，所有的2放到右边，1不考虑

func SortColors(nums []int) {
	n := len(nums)
	l := 0
	r := n - 1
	for i := 0; i < r; i++ {
		for i < r && nums[i] == 2 { //一直排到当前位置不可能==2，因为index从左边开始读，所以不允许左边有漏网之鱼
			nums[i], nums[r] = nums[r], nums[i]
			r--
		}
		if nums[i] == 0 {
			nums[i], nums[l] = nums[l], nums[i]
			l++
		}
	}
	fmt.Println(nums)
}

//我的两次循环来完成双指针遍历，时间无敌，空间爆表
//func SortColors(nums []int) {
//	n := len(nums)
//	l := 0
//	r := n - 1
//	for i := 0; i < n; i++ {
//		if nums[i] == 0 {
//			nums[i], nums[l]=nums[l],nums[i]
//			l++
//		}
//	}
//	for i := n - 1; i >= l; i-- {
//		if nums[i] == 2 {
//			nums[i], &nums[r]=nums[r],nums[i]
//			r--
//		}
//	}
//	fmt.Print(nums)
//}
