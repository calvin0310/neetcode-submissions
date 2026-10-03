type TrieNode struct {
	children [26]*TrieNode
	isWord bool
}

type WordDictionary struct {
    root *TrieNode
}

func Constructor() WordDictionary {
    return WordDictionary{
		root: &TrieNode{},
	}
}

func (this *WordDictionary) AddWord(word string)  {
    curr := this.root

	for i := 0; i < len(word); i++ {
		char := word[i]
		// 用數字當作 index，避免要用 26 層 switch 比對
		index := char - 'a'

		if curr.children[index] == nil {
			curr.children[index] = &TrieNode{}
		}

		curr = curr.children[index]
	}

	curr.isWord = true
}

func (this *WordDictionary) Search(word string) bool {
	// 定義遞迴函數，遞迴函數要先定義，避免編譯錯誤（ := 會先看右邊發現 dfs 而報錯）
	var dfs func(j int, curr *TrieNode) bool

	dfs = func(j int, curr *TrieNode) bool {
		for i := j; i < len(word); i++ {
			char := word[i]

			if char == '.' {
				// 不知道 index 是哪個，所以要 26 都找，最多只有兩個 '.'，所以最多只會找 26*26
				for _, child := range curr.children {
					if child != nil && dfs(i+1, child) {
						return true //dfs(i+1, child) 已經找到了
					}
				}

				return false // 26 個都找不到目標
			} else {
				index := char - 'a'
				if curr.children[index] == nil {
					return false
				}

				curr = curr.children[index] // 繼續找下一輪
			}
		}

		return curr.isWord
	}

	return dfs(0, this.root)
}
