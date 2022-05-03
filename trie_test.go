package trie

import "testing"

func Test_trie_Insert(t *testing.T) {
	trieNode := NewTrie()
	trieNode.Insert("abc")
}
