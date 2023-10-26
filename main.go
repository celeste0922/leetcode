package main

import (
	"fmt"
	"leetcode/Array"
)

func main() {
	fmt.Println(Array.TwoSum([]int{5, 1, 2, 3}, 3))
	fmt.Println(Array.MaxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
	fmt.Println(Array.ThreeSum([]int{-1, 0, 1, 2, -1, -4})) //-4,-1,-1,0,1,2
	fmt.Println(Array.ThreeSumClosest([]int{4, 0, 5, -5, 3, 3, 0, -4, -5}, -2))
}
