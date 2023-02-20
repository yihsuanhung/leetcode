package main

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */

type Trie struct {
	root *Node
}

type Node struct {
	char     string
	isEnd    bool
	children map[rune]*Node
}

func createNode(value string) *Node {
	return &Node{
		char:     value,
		isEnd:    false,
		children: make(map[rune]*Node),
	}
}

func Constructor() Trie {
	node := createNode("/")
	return Trie{root: node}
}

func (this *Trie) Insert(word string) {
	// 設定初始指針為root
	node := this.root
	// 走訪word中所有字
	for _, c := range word {
		// 檢查當前的字是否在當前指針，如果不在，就在當前指針的children新增node
		val, ok := node.children[c]
		if !ok {
			newNode := createNode(string(c))
			node.children[c] = newNode
		}
		// 更新指針
		node = val
	}
	// 走訪結束時，把當前指針的isEnd設定為true
	node.isEnd = true
}

func (this *Trie) Search(word string) bool {
	node := this.root
	for _, c := range word {
		val, ok := node.children[c]
		if !ok {
			return false
		}
		node = val
	}
	if node.isEnd == false {
		return false
	}
	return true

}

func (this *Trie) StartsWith(prefix string) bool {
	node := this.root
	for _, c := range prefix {
		_, ok := node.children[c]
		if !ok {
			return false
		}
	}
	return true
}

func main() {}
