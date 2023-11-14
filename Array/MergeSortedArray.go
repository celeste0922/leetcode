package Array

import "fmt"

//88. 合并两个有序数组
//给你两个按非递减顺序排列的整数数组nums1和nums2，另有两个整数m和n ，分别表示nums1和nums2中的元素数目。
//请你合并nums2到nums1中，使合并后的数组同样按 非递减顺序 排列。
//注意：最终，合并后数组不应由函数返回，而是存储在数组 nums1 中。
//为了应对这种情况，nums1的初始长度为m + n，其中前m个元素表示应合并的元素，后n个元素为0 ，应忽略。nums2的长度为n。
//==============================================================================================
//双指针=>倒序双指针

func Merge2(nums1 []int, m int, nums2 []int, n int) {
	index := len(nums1) - 1
	for n > 0 {
		if m > 0 && nums1[m-1] > nums2[n-1] {
			nums1[index] = nums1[m-1]
			index--
			m--
		} else {
			nums1[index] = nums2[n-1]
			n--
			index--
		}
	}
	fmt.Println(nums1)
}
