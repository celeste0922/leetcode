package Array

//78. 子集
//给你一个整数数组 nums ，数组中的元素 互不相同 。返回该数组所有可能的子集（幂集）。
//解集 不能 包含重复的子集。你可以按 任意顺序 返回解集。

func Subsets(nums []int) [][]int {
	result, subsetsPath = make([][]int, 0), make([]int, 0, len(nums))
	subsetsDfs(nums, 0)
	return result
}

var (
	result      [][]int
	subsetsPath []int
)

func subsetsDfs(nums []int, start int) {
	temp := make([]int, len(subsetsPath))
	copy(temp, subsetsPath)
	result = append(result, temp)
	for i := start; i < len(nums); i++ {
		subsetsPath = append(subsetsPath, nums[i])
		subsetsDfs(nums, i+1)
		subsetsPath = subsetsPath[:len(subsetsPath)-1]
	}
}
