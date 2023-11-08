package Array

//79. 单词搜索
//给定一个 m x n 二维字符网格 board 和一个字符串单词 word 。
//如果 word 存在于网格中，返回 true ；否则，返回 false 。
//单词必须按照字母顺序，通过相邻的单元格内的字母构成，
//其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。

func Exist(board [][]byte, word string) bool {
	existFound := false
	var existDfs func(board [][]byte, word string, i int, j int, k int)
	existDfs = func(board [][]byte, word string, i int, j int, k int) {
		if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) {
			return
		}
		if board[i][j] != word[k] || board[i][j] == '*' {
			return
		}
		if k == len(word)-1 {
			existFound = true
			return
		}
		temp := board[i][j]
		board[i][j] = '*'
		existDfs(board, word, i+1, j, k+1)
		existDfs(board, word, i, j+1, k+1)
		existDfs(board, word, i-1, j, k+1)
		existDfs(board, word, i, j-1, k+1)
		board[i][j] = temp
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			existDfs(board, word, i, j, 0)
		}
	}
	return existFound
}
