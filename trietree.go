package TrieTree

import (
	"container/list"
)

type Trie struct {
	charMap map[interface{}]*Trie
	isLeaf  false
}

func (t *Trie) Add(s string) bool {
	return false
}

func (t *Trie) Search(s string) bool {
	return false
}
