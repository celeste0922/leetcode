package Array

//216. 组合总和III  --Medium
//找出所有相加之和为n的k个数的组合，且满足下列条件：
//只使用数字1到9
//每个数字 最多使用一次
//返回所有可能的有效组合的列表 。该列表不能包含相同的组合两次，组合可以以任何顺序返回。
//=========================================================================
//递归回溯+剪枝

func CombinationSum3(k int, n int) [][]int {
	path3 = make([]int, 0, k)
	res3 = make([][]int, 0)
	dfs3(1, k, n)
	return res3
}

var (
	path3 []int
	res3  [][]int
)

func dfs3(index int, k int, target int) {
	if len(path3) == k {
		if target == 0 {
			temp := make([]int, len(path3))
			copy(temp, path3)
			res3 = append(res3, temp)
		}
		return
	}
	for i := index; i < 10; i++ {
		if target-i < 0 {
			break
		}
		path3 = append(path3, i)
		dfs3(i+1, k, target-i)
		path3 = path3[:len(path3)-1]
	}
}
