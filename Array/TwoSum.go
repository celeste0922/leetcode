package Array

//两数之和
//为要求的和与数组每个数之间的差建立map索引

func TwoSum(nums []int, target int) []int { //暴力
	myLenth := len(nums)
	myMap := map[int]int{}
	for i := 0; i < myLenth; i++ {
		//for j := i + 1; j < myLenth; j++ {
		//	if nums[i]+nums[j] == target {
		//		return []int{i, j}
		//	}
		//}

		if v, k := myMap[target-nums[i]]; k { //搜索table中是否存在target-num[i],存在则取下标并返回
			return []int{v, i}
		} else {
			myMap[nums[i]] = i //将num[i]作为key放入table便于搜索（索引）
		}

	}
	return []int{}
}

// 用空间换时间，创建table，对于每个x，先在表中搜索target-x,再把x插入table
func twoSum(nums []int, target int) []int {
	prevNums := map[int]int{}
	for i, num := range nums {
		targetNum := target - num
		targetNumIndex, ok := prevNums[targetNum]
		if ok {
			return []int{targetNumIndex, i}
		} else {
			prevNums[num] = i
		}
	}
	return []int{}
}

//func twoSum(nums []int, target int) []int {
//	hashTable := map[int]int{}
//	for i, x := range nums {
//		if p, ok := hashTable[target-x]; ok {
//			return []int{p, i}
//		}
//		hashTable[x] = i
//	}
//	return nil
//}
