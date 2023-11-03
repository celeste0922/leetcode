package Array

import "sort"

//56. 合并区间
//以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。
//请你合并所有重叠的区间，并返回 一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间 。

func Merge(intervals [][]int) [][]int {
	n := len(intervals)
	if n == 1 {
		return intervals
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	result := make([][]int, 0)
	for i := 0; i < n-1; i++ {
		if intervals[i][1] >= intervals[i+1][0] {
			intervals[i+1][0] = intervals[i][0]
			intervals[i+1][1] = max(intervals[i][1], intervals[i+1][1])
		} else {
			result = append(result, intervals[i])
		}
	}
	result = append(result, intervals[n-1])
	return result
}
