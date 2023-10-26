package Array

import "sort"

//最接近的三数之和
//给你一个长度为 n 的整数数组 nums和一个目标值target。请你从nums中选出三个整数，使它们的和与 target 最接近。
//返回这三个数的和。
//假定每组输入只存在恰好一个解。
//=============================================================
//思路类似于三数之和，优化重点在于在当前循环层跳过冗余的步骤（基于数组的排序对三数之和值的大小情况进行分类判断）
//======================================================================
//设s=nums[i]+nums[i+1]+nums[i+2]。如果 s>target，
//由于数组已经排序，后面无论怎么选，选出的三个数的和不会比s还小，所以不会找到比s更优的答案了。
//所以只要s>target，就可以直接 break 外层循环了。在 break 前判断s是否离target更近，如果更近，那么更新答案为s。
//设 s=nums[i]+nums[n−2]+nums[n−1]。如果 s<target，
//由于数组已经排序，nums[i]加上后面任意两个数都不超过 s，所以下面的双指针就不需要跑了，无法找到比s更优的答案。
//但是后面还有更大的 nums[i]，可能找到一个离 target
//更近的三数之和，所以还需要继续枚举,continue外层循环。在continue前判断s是否离target更近，
//如果更近，那么更新答案为s。
//如果 i>0且 nums[i]=nums[i−1]，那么 nums[i] 和后面数字相加的结果，必然在之前算出过，
//所以无需跑下面的双指针，直接continue外层循环。（可以放在循环开头判断。）

func ThreeSumClosest(nums []int, target int) int {
	n := len(nums)
	result := nums[0] + nums[1] + nums[n-1]
	sort.Ints(nums)
	for i := 0; i < n-2; i++ {
		if nums[i]+nums[i+1]+nums[i+2] > target {
			if getValue(nums[i]+nums[i+1]+nums[i+2]-target) < getValue(result-target) {
				result = nums[i] + nums[i+1] + nums[i+2]
			}
			break
		}
		if nums[n-1]+nums[n-2]+nums[n-3] < target {
			if getValue(nums[n-3]+nums[n-1]+nums[n-2]-target) < getValue(result-target) {
				result = nums[n-3] + nums[n-1] + nums[n-2]
			}
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		k := n - 1
		for j := i + 1; j < k; {
			if nums[n-1]+nums[n-2]+nums[i] < target {
				if getValue(nums[i]+nums[n-1]+nums[n-2]-target) < getValue(result-target) {
					result = nums[i] + nums[n-1] + nums[n-2]
				}
				break
			}
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
