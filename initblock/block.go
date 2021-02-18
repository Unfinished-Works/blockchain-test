package initblock

import (
	"encoding/hex"

	"github.com/Unfinished-Works/blockchain-test/blockchain"
	"github.com/Unfinished-Works/blockchain-test/blockstore"
	"github.com/Unfinished-Works/blockchain-test/pow"
	"google.golang.org/protobuf/proto"
)

//InitBlock :
func InitBlock() {
	fs := blockstore.InitFileStore("blocks")
	if _, err := fs.Get("ace8b9baeed95ad0fd8b58e010a3e99d553af964779032e35daf6ea330a93a6d"); err == nil {
		return
	}
	GenesisBlock := new(blockchain.Block)
	GenesisBlock.Nonce = 0
	GenesisBlock.Header = new(blockchain.Header)
	GenesisBlock.Header.Ref = make([]byte, 32)
	GenesisBlock.Header.ID = 0
	GenesisBlock.Header.Difficulty = 0
	GenesisBlock.Header.Timestamp = 1613633682931145500
	GenesisBlock.HeaderHash = GenesisBlock.Header.Hash()
	GenesisBlock.Tx = []*blockchain.Transaction{}
	GenesisBlock.SealedBy = ""
	GenesisBlock.Data = []byte{}
	block, _ := proto.Marshal(GenesisBlock)
	hash := pow.HashFunc(block, GenesisBlock.GetNonce())
	f, err := fs.Store(hex.EncodeToString(hash))
	if err != nil {
		panic(err)
	}
	f.Write(block)
	f.Sync()
	f.Close()
}
