package main

import "fmt"

type Trie struct {
	Children [26]*Trie
	IsEnd    bool
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	node := this
	for _, ch := range word {
		ch -= 'a'
		if node.Children[ch] == nil {
			node.Children[ch] = &Trie{}
		}
		node = node.Children[ch]
	}
	node.IsEnd = true
}
func (this *Trie) searchPrefix(prefix string) *Trie {
	node := this
	for _, ch := range prefix {
		ch -= 'a'
		if node.Children[ch] == nil {
			return nil
		}
		node = node.Children[ch]
	}
	return node
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	node := this.searchPrefix(word)
	return node != nil && node.IsEnd
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	return this.searchPrefix(prefix) != nil
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
func main() {
	t := new(Trie)
	t.Insert("apple")
	t.Search("apple")
	t.StartsWith("app")
	t.Insert("app")
	t.Search("app")
	fmt.Println(t.Search("apple"),
		t.StartsWith("app"),
		t.Search("app"))
}
