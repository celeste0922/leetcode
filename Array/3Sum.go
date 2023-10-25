package Array

import "sort"

// 三数之和
// 从小到大排序
// 第一个数作为第一层循环
// 第二个数和第三个数作为第二层循环的双指针
// 第二层遇到相同的下个数自动跳过避免重复

func ThreeSum(nums []int) [][]int {
	i := 0

	index := 0
	result := [][]int{}
	sums := []int{}
	sort.Ints(nums)
	for ; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		j := i + 1
		k := len(nums) - 1
		for j < k {
			if j-1 > i && nums[j] == nums[j-1] {
				j++
				continue
			}
			if k+1 < len(nums) && nums[k] == nums[k+1] {
				k--
				continue
			}
			sum := nums[i] + nums[j] + nums[k]
			if sum == 0 {
				sums = []int{nums[i], nums[j], nums[k]}
				result = append(result, sums)
				index++
				j++
			} else if sum > 0 {
				k--
			} else {
				j++
			}
		}
	}
	return result
}
