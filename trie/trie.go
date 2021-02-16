package trie

import "github.com/Unfinished-Works/blockchain-test/blockchain"

//Trie : tree node based BlockStore
type Trie struct {
	isroot bool
	depth  int
	Hash   string
	Block  *blockchain.Block
	Childs []*Trie
}

func (x *Trie) rawInsert(new *Trie) {
	new.depth = x.depth + 1
	x.Childs = append(x.Childs, new)
}

//Insert : Add a node to the Triechain
func (x *Trie) Insert(ref string, new *Trie) bool {
	parent, ok := x.findnode(ref)
	if !ok {
		return false
	}
	parent.rawInsert(new)
	return true
}

func (x *Trie) findnode(hash string) (*Trie, bool) {
	if x.Hash == hash {
		return x, true
	}
	if len(x.Childs) == 0 {
		return nil, false
	}
	for i := range x.Childs {
		node, ok := x.Childs[i].findnode(hash)
		if ok {
			return node, true
		}
	}
	return nil, false
}
