package Array

import "sort"

//最接近的三数之和
//给你一个长度为 n 的整数数组 nums 和 一个目标值 target。请你从 nums 中选出三个整数，使它们的和与 target 最接近。
//返回这三个数的和。
//假定每组输入只存在恰好一个解。
//=============================================================
//思路类似于三数之和，优化重点在于在当前循环层跳过冗余的步骤（基于数组的排序对三数之和值的大小情况进行分类判断）

func ThreeSumClosest(nums []int, target int) int {
	n := len(nums)
	result := 100
	sort.Ints(nums)
	for i := 0; i < n-2; i++ {
		if nums[i]+nums[i+1]+nums[i+2] > target {
			if getValue(nums[i]+nums[i+1]+nums[i+2]-target) < getValue(result-target) {
				result = nums[i] + nums[i+1] + nums[i+2]
			}
			break
		}
		if nums[n-1]+nums[n-2]+nums[n-3] < target {
			result = nums[n-1] + nums[n-2] + nums[n-3]
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		k := n - 1
		for j := i + 1; j < k; {
			if j > i+1 && nums[j] == nums[j-1] {
				j++
				continue
			}
			if k < n-1 && nums[k] == nums[k+1] {
				k--
				continue
			}
			flag := nums[i] + nums[j] + nums[k] - target
			if flag == 0 {
				return nums[i] + nums[j] + nums[k]
			}
			if getValue(flag) < getValue(result-target) {
				result = nums[i] + nums[j] + nums[k]
			} else if flag > 0 {
				k--
			} else {
				j++
			}
		}
	}
	return result
}

func getValue(t int) int {
	if t >= 0 {
		return t
	} else {
		return -t
	}
}
