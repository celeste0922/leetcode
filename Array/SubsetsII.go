package Array

import (
	"sort"
)

// 90. 子集II
// 给你一个整数数组nums ，其中可能包含重复元素，请你返回该数组所有可能的子集（幂集）。
// 解集不能包含重复的子集。返回的解集中，子集可以按任意顺序排列。
// ====================================================
//go数组之间的赋值要注意是引用赋值还是常量赋值

var (
	res2  [][]int
	path2 []int
)

func SubsetsWithDup(nums []int) [][]int {
	res2 = make([][]int, 0)
	path2 = make([]int, 0, len(nums))
	sort.Ints(nums)
	dfs2(nums, 0)
	return res2
}

func dfs2(nums []int, index int) {
	temp := make([]int, len(path2))
	copy(temp, path2) //一定要记得使用copy，否则底层赋值引用，res的值会随path变化
	res2 = append(res2, temp)
	for i := index; i < len(nums); i++ {
		if i > index && nums[i] == nums[i-1] {
			continue
		}
		//res2 = append(res2, path2)
		path2 = append(path2, nums[i])
		dfs2(nums, i+1)
		path2 = path2[:len(path2)-1]
	}
}
