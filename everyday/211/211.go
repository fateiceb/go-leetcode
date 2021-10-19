package main

type TrieNode struct {
	Children [26]*TrieNode
	Isend    bool
}

func (t *TrieNode) Insert(word string) {
	node := t
	for _, ch := range word {
		ch = ch - 'a'
		if node.Children[ch] == nil {
			node = &TrieNode{}
		}
		node = node.Children[ch]
	}
	node.Isend = true
}

type WordDictionary struct {
	trieRoot *TrieNode
}

func Constructor() WordDictionary {
	return WordDictionary{&TrieNode{}}
}

func (this *WordDictionary) AddWord(word string) {
	this.trieRoot.Insert(word)
}

func (this *WordDictionary) Search(word string) bool {
	var dfs func(int, *TrieNode) bool
	dfs = func(index int, node *TrieNode) bool {
		if index == len(word) {
			return node.Isend
		}
		ch := word[index]
		if ch != '.' {
			child := node.Children[ch-'a']
			if child != nil && dfs(index+1, child) {
				return true
			}
		} else {
			for i := range node.Children {
				child := node.Children[i]
				if child != nil && dfs(index+1, child) {
					return true
				}
			}
		}
		return false
	}
	return dfs(0, this.trieRoot)
}
func main() {

}
