package main

type WordDictionary struct {
	Children [26]*WordDictionary
	isEnd    bool
}

func Constructor() WordDictionary {
	return WordDictionary{[26]*WordDictionary{}, false}
}

func (this *WordDictionary) AddWord(word string) {
	cur := this
	for _, w := range word {
		if cur.Children[w-'a'] == nil {
			tmp := Constructor()
			cur.Children[w-'a'] = &tmp
		}
		cur = this.Children[w-'a']
	}
	cur.isEnd = true
}

func (this *WordDictionary) Search(word string) bool {
	cur := this
	for i, w := range word {
		// handle case: "."
		if word[i] == '.' {
			for _, item := range cur.Children {
				if item != nil && item.Search(word[i+1:]) {
					return true
				}
			}

			return false
		}

		if cur.Children[w-'a'] == nil {
			return false
		}
		cur = cur.Children[w-'a']
	}
	return cur.isEnd
}

func main() {

}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
