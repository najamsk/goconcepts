package main

import "fmt"

const TotalAlphabets = 26

type Node struct {
	children [TotalAlphabets]*Node
	isEnd    bool
}

//hold root node
type Trie struct {
	root *Node
}

func InitTrie() *Trie {
	res := &Trie{root: &Node{}}
	return res
}

func (t *Trie) Insert(w string) {
	wordLen := len(w)
	cNode := t.root
	for i := 0; i < wordLen; i++ {
		charIndex := w[i] - 'a'
		if cNode.children[charIndex] == nil {
			cNode.children[charIndex] = &Node{}
		}
		cNode = cNode.children[charIndex]
	}
	cNode.isEnd = true
}

func (t *Trie) Search(w string) bool {
	wordLen := len(w)
	cNode := t.root
	for i := 0; i < wordLen; i++ {
		charIndex := w[i] - 'a'
		if cNode.children[charIndex] == nil {
			return false
		}
		cNode = cNode.children[charIndex]
	}
	if cNode.isEnd == true {
		return true
	}

	return false
}

func main() {
	fmt.Println("hello")
	t := InitTrie()
	fmt.Println(t.root)

	words := []string{"hello", "world", "najam"}
	for _, v := range words {
		t.Insert(v)
	}
	for _, v := range words {
		fmt.Println(t.Search(v))
	}
	fmt.Println(t.Search("najaf"))

}
