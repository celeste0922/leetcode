package Array

import "fmt"

//127. 单词接龙   --Hard
//字典 wordList 中从单词 beginWord 和 endWord 的 转换序列 是一个按下述规格形成的序列 beginWord -> s1 -> s2 -> ... -> sk：
//每一对相邻的单词只差一个字母。
//对于 1 <= i <= k 时，每个 si 都在 wordList 中。注意， beginWord 不需要在 wordList 中。
//sk == endWord
//给你两个单词 beginWord 和 endWord 和一个字典 wordList ，返回 从 beginWord 到 endWord 的 最短转换序列 中的 单词数目 。如果不存在这样的转换序列，返回 0 。
//===========================================================================================
//单向bfs

func LadderLength(beginWord string, endWord string, wordList []string) int {
	//将wordList放入Map，标志是否已经过节点
	wordMap := map[string]bool{}
	for _, w := range wordList {
		wordMap[w] = true
	}
	path := []string{beginWord}
	pathLength := 1
	for len(path) != 0 {
		pathSize := len(path)
		for i := 0; i < pathSize; i++ {
			word := path[0]
			if word == endWord {
				return pathLength
			}
			path = path[1:]
			for j := 0; j < len(word); j++ {
				for a := 'a'; a <= 'z'; a++ {
					newWord := word[:j] + string(a) + word[j+1:]
					if wordMap[newWord] == true {
						fmt.Println(newWord)
						path = append(path, newWord)
						wordMap[newWord] = false
					}
				}
			}
		}
		pathLength++
	}
	return 0
}

// ====================================双向bfs======================================
func ladderLength(beginWord string, endWord string, wordList []string) int {
	// wordList 转 map
	wordMap := map[string]struct{}{}
	for _, w := range wordList {
		wordMap[w] = struct{}{}
	}
	if _, ok := wordMap[endWord]; !ok {
		return 0
	}
	visited := map[string]struct{}{}
	beginVisited := map[string]struct{}{}
	beginVisited[beginWord] = struct{}{}
	endVisited := map[string]struct{}{}
	endVisited[endWord] = struct{}{}
	wordLen := len(beginWord)
	step := 1
	for len(beginVisited) != 0 && len(endVisited) != 0 {
		if len(beginVisited) > len(endVisited) {
			beginVisited, endVisited = endVisited, beginVisited
		}
		fmt.Println(beginVisited, step, endVisited)
		nextVisted := map[string]struct{}{}
		for word := range beginVisited {
			for posi := 0; posi < wordLen; posi++ {
				var bt byte
				for bt = 'a'; bt <= 'z'; bt++ {
					if bt == word[posi] {
						continue
					}
					wordBytes := []byte(word)
					wordBytes[posi] = bt
					newWord := string(wordBytes)
					if _, ok := wordMap[newWord]; ok {
						if _, ok := endVisited[newWord]; ok {
							return step + 1
						}
						if _, ok := visited[newWord]; !ok {
							visited[newWord] = struct{}{}
							nextVisted[newWord] = struct{}{}
						}
					}
				}
			}
		}
		beginVisited = nextVisted
		step++
	}

	return 0
}
