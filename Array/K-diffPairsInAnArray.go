package Array

import (
	"sort"
)

//532. 数组中的k-diff数对   --Medium
//给你一个整数数组 nums 和一个整数 k，请你在数组中找出 不同的 k-diff 数对，并返回不同的 k-diff 数对 的数目。
//k-diff 数对定义为一个整数对 (nums[i], nums[j]) ，并满足下述全部条件：
//0 <= i, j < nums.length
//i != j
//|nums[i] - nums[j]| == k
//注意，|val| 表示 val 的绝对值。
//=============================
//哈希表或双指针
//核心思想都是num[fast]-num[slow]==k
//关键优化点=>遍历到的每个num[fast]只有一个num[slow]与之对应

func FindPairs(nums []int, k int) int {
	n := len(nums)
	sort.Ints(nums)
	res := 0
	for slow, fast := 0, 1; fast < n; {
		for fast < n-1 && nums[fast] == nums[fast+1] {
			fast++
		}
		if slow == fast {
			fast++
			continue
		}
		temp := nums[fast] - nums[slow]
		if temp == k {
			res++
		}
		if temp > k {
			slow++
			continue
		}
		fast++
	}
	return res
}

//=====================优雅=========================
//关键优化点=>遍历到的每个num[fast]只有一个num[slow]与之对应
//func findPairs(nums []int, k int) (ans int) {
//	sort.Ints(nums)
//	y, n := 0, len(nums)
//	for x, num := range nums {
//		if x == 0 || num != nums[x-1] {
//			for y < n && (nums[y] < num+k || y <= x) {
//				y++
//			}
//			if y < n && nums[y] == num+k {
//				ans++
//			}
//		}
//	}
//	return
//}
