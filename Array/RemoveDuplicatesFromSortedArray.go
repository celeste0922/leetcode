package Array

//删除有序数组中的重复项
//给你一个非严格递增排列的数组 nums ，请你原地删除重复出现的元素，使每个元素只出现一次，
//返回删除后数组的新长度。元素的相对顺序应该保持一致 。然后返回 nums 中唯一元素的个数。
//考虑 nums 的唯一元素的数量为 k ，你需要做以下事情确保你的题解可以被通过：
//更改数组 nums ，使 nums 的前 k 个元素包含唯一元素，并按照它们最初在 nums 中出现的顺序排列。nums 的其余元素与 nums 的大小不重要。
//返回 k 。
//判题标准:
//系统会用下面的代码来测试你的题解:
//int[] nums = [...]; // 输入数组
//int[] expectedNums = [...]; // 长度正确的期望答案
//int k = removeDuplicates(nums); // 调用
//assert k == expectedNums.length;
//for (int i = 0; i < k; i++) {
//assert nums[i] == expectedNums[i];
//}
//如果所有断言都通过，那么您的题解将被通过。
//==================================================================================
//经典双指针
//一个指针指向存储位置，一个指针指向读取位置

func RemoveDuplicates(nums []int) int {
	n := len(nums)
	p := 0
	q := 0
	if n == 0 {
		return 0
	}
	for q < n { //1,1,2
		if nums[p] == nums[q] {
			q++
		} else {
			p++
			nums[p] = nums[q]
		}
	}
	return p + 1
}

//=================官方题解==优雅=======================
//func removeDuplicates(nums []int) int {
//	n := len(nums)
//	if n == 0 {
//		return 0
//	}
//	slow := 1
//	for fast := 1; fast < n; fast++ {
//		if nums[fast] != nums[fast-1] {
//			nums[slow] = nums[fast]
//			slow++
//		}
//	}
//	return slow
//}
