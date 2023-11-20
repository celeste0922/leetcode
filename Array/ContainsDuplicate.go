package Array

//217. 存在重复元素   --Easy
//给你一个整数数组nums。如果任一值在数组中出现至少两次 ，返回true;
//如果数组中每个元素互不相同，返回false。
//====================================================
//创建一个哈希表，然后从左往右遍历数组。
//检测哈希表中是否已存在当前字符，若存在，直接返回结果，
//若不存在，将当前字符加入哈希表，供后续判断使用即可。

func ContainsDuplicate(nums []int) bool {
	set := map[int]struct{}{}
	for _, v := range nums {
		if _, has := set[v]; has { //叼叼叼
			return true
		}
		set[v] = struct{}{}
	}
	return false
}
