package trie

type trie struct {
	root *node
	size int
}

type node struct {
	val      rune
	last     bool
	children map[rune]*node
	parent   *node
	depth    int
}

func New() *trie {
	return &trie{
		root: &node{
			children: make(map[rune]*node),
			depth:    0,
		},
		size: 0,
	}
}

func (t *trie) Add(s string) {
	nd := t.root
	for i, r := range s {
		isLast := i == len(s)-1
		nd2, ok := nd.children[r]
		if ok {
			if isLast && !nd2.last {
				nd2.last = true
				t.size++
				return
			}
			nd = nd2
			continue
		}
		nd2 = &node{
			val:      r,
			children: make(map[rune]*node),
			parent:   nd,
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

func (t *trie) Find(s string) bool {
	_, exists := t.findNode(s)
	return exists
}

func (t *trie) findNode(s string) (lastNode *node, exists bool) {
	nd := t.root
	for i, r := range s {
		isLast := i == len(s)-1
		nd2, ok := nd.children[r]
		if !ok {
			return nil, false
		}
		if isLast && nd2.last {
			lastNode = nd2
			exists = true
			return
		}
		nd = nd2
	}
	return nil, false
}

func (t *trie) Remove(s string) {
	//find
	lastNode, exists := t.findNode(s)
	if !exists {
		return
	}
	t.size--
	if len(lastNode.children) > 0 {
		lastNode.last = false
		return
	}
	//remove
	parent := lastNode.parent
	r := lastNode.val
	for len(parent.children) <= 1 {
		delete(parent.children, r)
		if parent.last {
			return
		}
		parent = parent.parent
		r = parent.val
	}
}

func (t *trie) Size() int {
	return t.size
}
