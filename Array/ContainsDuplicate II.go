package Array

//219. 存在重复元素II   --Easy
//给你一个整数数组nums和一个整数k,判断数组中是否存在两个不同的索引i和j,
//满足nums[i] == nums[j]且abs(i - j) <= k.如果存在,返回 true;
//否则,返回 false.
//=======================================================
//滑动窗口
//叼叼叼

func ContainsNearbyDuplicate(nums []int, k int) bool {
	set := map[int]struct{}{}
	for i, num := range nums {
		//当在数组中读取到>k位置时，每读取一位，删除哈希表的一位，实现滑动窗口
		if i > k {
			delete(set, nums[i-k-1])
		}
		if _, ok := set[num]; ok {
			return true
		}
		set[num] = struct{}{}
	}
	return false
}

//func ContainsNearbyDuplicate(nums []int, k int) bool {
//	set := make(map[int]int)
//	for key, v := range nums {
//		if key1, has := set[v]; has {
//			if abs(key-key1) <= k {
//				return true
//			}
//		}
//		set[v] = key
//	}
//	return false
//}
//func abs(a int) int {
//	if a > 0 {
//		return a
//	} else {
//		return -a
//	}
//}
