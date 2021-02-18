package blockchain

import (
	"io"

	"github.com/Unfinished-Works/blockchain-test/blockstore"
	"google.golang.org/protobuf/proto"
)

/*
//Trie : tree node based BlockStore
type Trie struct {
	IsRoot bool
	Depth  int
	Hash   string
	Block  *blockchain.Block
	Childs []*Trie
}*/

//NewTrie : Init New Trie
func NewTrie() *Trie {
	//GenesisBlock : first Block
	var fs = blockstore.InitFileStore("blocks")
	var GenesisBlock = func() *Block {
		f, err := fs.Get(GenesisBlockHash)
		if err != nil {
			panic(err)
		}
		b, err := io.ReadAll(f)
		if err != nil {
			panic(err)
		}
		gblock := new(Block)
		gblock.Header = new(Header)
		err = proto.Unmarshal(b, gblock)
		if err != nil {
			panic(err)
		}
		return gblock
	}()
	return &Trie{
		IsRoot: true,
		Depth:  0,
		Hash:   GenesisBlockHash,
		Block:  GenesisBlock,
		Childs: nil,
	}
}

func (x *Trie) rawInsert(new *Trie) {
	new.Depth = x.Depth + 1
	x.Childs = append(x.Childs, new)
}

//Insert : Add a node to the Triechain
func (x *Trie) Insert(ref string, new *Trie) bool {
	parent, ok := x.Findnode(ref)
	if !ok {
		return false
	}
	parent.rawInsert(new)
	return true
}

//Findnode : Search for hash in Triechain
func (x *Trie) Findnode(hash string) (*Trie, bool) {
	if x.Hash == hash {
		return x, true
	}
	if len(x.Childs) == 0 {
		return nil, false
	}
	for i := range x.Childs {
		node, ok := x.Childs[i].Findnode(hash)
		if ok {
			return node, true
		}
	}
	return nil, false
}
