package trie

type trie struct {
	root *node
	size int
}

type node struct {
	val      rune
	last     bool
	children map[rune]*node
	depth    int
}

func NewTrie() *trie {
	return &trie{
		root: &node{
			children: make(map[rune]*node),
			depth:    0,
		},
		size: 0,
	}
}

func (t *trie) Insert(s string) {
	nd := t.root
	for i, r := range s {
		isLast := i == len(s)-1
		nd2, ok := nd.children[r]
		if ok {
			if !nd2.last && isLast {
				nd2.last = true
				t.size++
				return
			}
			continue
		}
		nd2 = &node{
			val:      r,
			children: make(map[rune]*node),
			depth:    nd.depth + 1,
		}
		nd.children[r] = nd2
		if isLast {
			nd2.last = true
			t.size++
			return
		}
		nd = nd2
	}
}
