package blockchain

import (
	"encoding/binary"

	"github.com/zeebo/blake3"
)

//GenesisBlockHash : first block
const GenesisBlockHash = "ace8b9baeed95ad0fd8b58e010a3e99d553af964779032e35daf6ea330a93a6d"

//Hash : Compute the blake3 hash.
func (x *Header) Hash() []byte {
	buf := make([]byte, 24)
	binary.BigEndian.PutUint64(buf[:8], x.Difficulty)
	binary.BigEndian.PutUint64(buf[8:16], x.ID)
	binary.BigEndian.PutUint64(buf[16:24], x.Timestamp)
	h := blake3.New()
	h.Write(x.Ref)
	h.Write(buf)
	return h.Sum(nil)
}
