package Array

import (
	"sort"
)

//四数之和
//给你一个由n个整数组成的数组 nums ，和一个目标值 target 。
//请你找出并返回满足下述全部条件且不重复的四元组 [nums[a], nums[b], nums[c], nums[d]]
//（若两个四元组元素一一对应，则认为两个四元组重复）：
//0 <= a, b, c, d < n
//a、b、c 和 d 互不相同
//nums[a] + nums[b] + nums[c] + nums[d] == target
//你可以按任意顺序返回答案 。
//=================================================
//n数之和通用解法，排序+双指针
//排序+固定枚举前n-2个数字，双指针搜索最后两个数字。
//===============================================================
//快捷版：直接调用三数之和并在外层嵌套一层循环（记得=剪枝=不然会溢出）
//疯狂剪枝==考虑边界情况

func FourSum(nums []int, target int) [][]int {
	n := len(nums)
	sort.Ints(nums)
	result := [][]int{}
	for p := 0; p < n-3; p++ {
		if nums[p]+nums[p+1]+nums[p+2]+nums[p+3] > target {
			break
		}
		if nums[n-1]+nums[n-2]+nums[n-3]+nums[n-3] < target {
			break
		}
		if nums[p]+nums[n-1]+nums[n-2]+nums[n-3] < target {
			continue
		}
		if p > 0 && nums[p] == nums[p-1] {
			continue
		}
		for q := p + 1; q < n-2; q++ {
			if nums[p]+nums[q]+nums[q+1]+nums[q+2] > target {
				break
			}
			if nums[p]+nums[q]+nums[n-1]+nums[n-2] < target {
				continue
			}
			if q > p+1 && nums[q] == nums[q-1] {
				continue
			}
			k := n - 1
			for j := q + 1; j < k; {
				if j > q+1 && nums[j] == nums[j-1] {
					j++
					continue
				}
				if k < n-1 && nums[k] == nums[k+1] {
					k--
					continue
				}
				if nums[p]+nums[q]+nums[j]+nums[k] == target {
					result = append(result, []int{nums[p], nums[q], nums[j], nums[k]})
					j++
				} else if nums[p]+nums[q]+nums[j]+nums[k] < target {
					j++
				} else {
					k--
				}
			}
		}
	}
	return result
}
