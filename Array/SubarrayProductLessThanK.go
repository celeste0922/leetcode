package Array

import "fmt"

//713. 乘积小于 K 的子数组  --Medium
//给你一个整数数组nums和一个整数k ，请你返回子数组内所有元素的乘积严格小于k的连续子数组的数目。
//==============================================================================
//滑动窗口=>遍历数组，途中遇到乘积temp>=k的情况直接for循环往前滑

func NumSubarrayProductLessThanK(nums []int, k int) int {
	n := len(nums)
	l := 0
	r := 0
	temp := 1
	res := 0
	for l < n {

		for temp >= k {
			if r < l {
				r = l
			}
			temp /= nums[l]
			l++
		}

		if temp < k {
			res++
			if r < n-1 {
				temp *= nums[r]
				fmt.Println(l, r, res, temp)
				r++
			} else {
				temp /= nums[l]
				fmt.Println(l, r, res, temp)
				l++
			}
			continue
		}

	}
	return res
}

//func numSubarrayProductLessThanK(nums []int, k int) (ans int) {
//	prod, i := 1, 0
//	for j, num := range nums {
//		prod *= num
//		for ; i <= j && prod >= k; i++ {
//			prod /= nums[i]
//		}
//		ans += j - i + 1
//	}
//	return
//}
