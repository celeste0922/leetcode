package Array

//713. 乘积小于 K 的子数组  --Medium
//给你一个整数数组nums和一个整数k ，请你返回子数组内所有元素的乘积严格小于k的连续子数组的数目。
//==============================================================================
//滑动窗口=>遍历数组，途中遇到乘积temp>=k的情况直接for循环往前滑

func NumSubarrayProductLessThanK(nums []int, k int) (ans int) {
	prod, i := 1, 0
	for j, num := range nums {
		prod *= num
		for ; i <= j && prod >= k; i++ {
			prod /= nums[i]
		}
		ans += j - i + 1
		//fmt.Println(j, i, ans)
	}
	return
}
