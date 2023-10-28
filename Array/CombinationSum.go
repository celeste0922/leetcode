package Array

import "sort"

// 39. 组合总和
//给你一个 无重复元素 的整数数组 candidates 和一个目标整数 target ，
//找出 candidates 中可以使数字和为目标数 target 的所有不同组合 ，并以列表形式返回。你可以按任意顺序返回这些组合。
//candidates 中的同一个数字可以无限制重复被选取 。如果至少一个数字的被选数量不同，则两种组合是不同的。
//对于给定的输入，保证和为 target 的不同组合数少于 150 个。
//===================================================================
//对于这类寻找所有可行解的题，我们都可以尝试用「搜索回溯」的方法来解决。
//回到本题，我们定义递归函数 dfs(target,combine,idx)表示当前在 candidates数组的第 idx 位，还剩 target要组合，已经组合的列表为 combine。
//递归的终止条件为 target≤0 或者 candidates数组被全部用完。那么在当前的函数中，每次我们可以选择跳过不用第 idx个数，即执行 dfs(target,combine,idx+1)。
//也可以选择使用第 idx个数，即执行 dfs(target−candidates[idx],combine,idx), 注意到每个数字可以被无限制重复选取，因此搜索的下标仍为 idx。
//
//更形象化地说，如果我们将整个搜索过程用一个树来表达，每次的搜索都会延伸出两个分叉，直到递归的终止条件，这样我们就能不重复且不遗漏地找到所有可行解。
//===================================
//剪枝优化

func CombinationSum(candidates []int, target int) [][]int {
	res, path = make([][]int, 0), make([]int, 0, len(candidates)) //切片初始化，（类型，长度，容量）内存问题？
	sort.Ints(candidates)                                         //排序：为剪枝优化做准备
	dfs(candidates, target, 0)
	return res
}

var (
	res  [][]int //结果集
	path []int   //路径记录
)

func dfs(candidates []int, target int, index int) {
	if target == 0 { //当传入的target=0说明前面路径上的值之和已经符合要求
		//新建一个切片并存入res，避免因底层内存共享而产生值的变化
		temp := make([]int, len(path))
		copy(temp, path)

		res = append(res, temp)
		return
	}
	for ; index < len(candidates); index++ { //index，按顺序访问节点，当前层访问过的节点不再访问避免重复
		if candidates[index] > target { //第一个节点大于0则剩余节点都不符合要求-直接结束循环（已排序）
			break
		}
		path = append(path, candidates[index])           //添加当前层节点路径（足迹）以便进入下一层节点
		dfs(candidates, target-candidates[index], index) //进入下一层节点
		path = path[:len(path)-1]                        //下一层节点已回溯，删除路径
	}
}
