package Array

import "sort"

// 三数之和
//给你一个整数数组 nums ，判断是否存在三元组 [nums[i], nums[j], nums[k]] 满足 i != j、i != k 且 j != k ，同时还满足 nums[i] + nums[j] + nums[k] == 0 。请
//你返回所有和为 0 且不重复的三元组。
//注意：答案中不可以包含重复的三元组。
//===============================================
// 从小到大排序
// 第一个数作为第一层循环
// 第二个数和第三个数作为第二层循环的双指针
// 第二层遇到相同的下个数自动跳过避免重复
//===============================================
//设 s = nums[first] + nums[first+1] + nums[first+2]，如果 s > 0，由于数组已经排序，后面无论怎么选，选出的三个数的和不会比 s 还小，所以只要 s > 0 就可以直接 break 外层循环了。
//
//如果 nums[first] + nums[n-2] + nums[n-1] < 0，由于数组已经排序，nums[first] 加上后面任意两个数都是小于 0 的，所以下面的双指针就不需要跑了。但是后面可能有更大的 nums[first]，所以还需要继续枚举，continue 外层循环。
//
//加上这两个优化后，就可以让代码击败接近 100%

func ThreeSum(nums []int) [][]int {
	i := 0
	n := len(nums)
	index := 0
	result := [][]int{}
	sums := []int{}
	sort.Ints(nums)
	for ; i < n-2; i++ {
		if nums[i]+nums[i+1]+nums[i+2] > 0 {
			break
		}
		if nums[i]+nums[n-2]+nums[n-1] < 0 {
			continue
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		j := i + 1
		k := n - 1
		for j < k {
			if j-1 > i && nums[j] == nums[j-1] {
				j++
				continue
			}
			if k+1 < n && nums[k] == nums[k+1] {
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
