package main

import (
	"fmt"
	"leetcode/Array"
)

func main() {

	//===================Array======================
	fmt.Print("两数之和===>")
	fmt.Println(Array.TwoSum([]int{5, 1, 2, 3}, 3))
	fmt.Print("最大容量数之和===>")
	fmt.Println(Array.MaxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
	fmt.Print("三数之和===>")
	fmt.Println(Array.ThreeSum([]int{-1, 0, 1, 2, -1, -4})) //-4,-1,-1,0,1,2
	fmt.Print("最接近的三数之和===>")
	fmt.Println(Array.ThreeSumClosest([]int{4, 0, 5, -5, 3, 3, 0, -4, -5}, -2))
	fmt.Print("四数之和===>")
	fmt.Println(Array.FourSum([]int{1, 0, -1, 0, -2, 2}, 0))
	fmt.Print("删除有序数组中的重复项===>")
	fmt.Println(Array.RemoveDuplicates([]int{1, 1, 2}))
	fmt.Print("移除元素===>")
	fmt.Println(Array.RemoveElement([]int{3, 2, 2, 3}, 3))
	fmt.Print("组合总和===>")
	fmt.Println(Array.CombinationSum([]int{2, 3, 5}, 8))
}
