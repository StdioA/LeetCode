package main

import "fmt"

type Trie struct {
    End bool
    Tree []*Trie
}

func newTrie() *Trie {
	trie := Trie{}
	trie.Tree = make([]*Trie, 26)
	return &trie
}

/** Initialize your data structure here. */
func Constructor() Trie {
    return *newTrie()
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string)  {
	node := this
    for _, r := range word {
		rn := r - 'a'
		n:= node.Tree[rn]
		if n == nil {
			node.Tree[rn] = newTrie()
			n = node.Tree[rn]
		}
		node = n
	}
	node.End = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	node := this
	for _, r := range word {
		node = node.Tree[r-'a']
		if node == nil {
			return false
		}
	}
	return node != nil && node.End
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	node := this
	for _, r := range prefix {
		node = node.Tree[r-'a']
		if node == nil {
			return false
		}
	}
	return node != nil
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
func main() {
	trie := Constructor()

	trie.Insert("apple")
	fmt.Println(trie.Search("apple"))   // returns true
	fmt.Println(trie.Search("app"))     // returns false
	fmt.Println(trie.StartsWith("app")) // returns true
	trie.Insert("app")   
	fmt.Println(trie.Search("app"))     // returns true
}
