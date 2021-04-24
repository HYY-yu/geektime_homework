package 第四周

type Trie struct {
	IsEnd    bool
	children map[uint8]*Trie
}

/** Initialize your data structure here. */
func Constructor() Trie {
	t := Trie{}
	t.children = make(map[uint8]*Trie)
	return t
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	p := this
	for i := 0; i < len(word); i++ {
		s := word[i]
		if v, ok := p.children[s]; ok {
			p = v
		} else {
			n := &Trie{}
			n.children = make(map[uint8]*Trie)
			p.children[s] = n
			p = n
		}
	}
	p.IsEnd = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	p := this
	for i := 0; i < len(word); i++ {
		s := word[i]
		if v, ok := p.children[s]; ok {
			p = v
		} else {
			return false
		}
	}
	return p.IsEnd
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	p := this
	for i := 0; i < len(prefix); i++ {
		s := prefix[i]
		if v, ok := p.children[s]; ok {
			p = v
		} else {
			return false
		}
	}
	return true
}
