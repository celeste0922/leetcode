package Array

//57. 插入区间
//给你一个无重叠的，按照区间起始端点排序的区间列表。
//在列表中插入一个新的区间，你需要确保列表中的区间仍然有序且不重叠（如果有必要的话，可以合并区间）。
//==============================================================================
//直接套用上一集！

func Insert(intervals [][]int, newInterval []int) [][]int {
	intervals = append(intervals, newInterval)
	return Merge(intervals)
}
