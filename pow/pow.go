package pow

import (
	"encoding/binary"

	"github.com/zeebo/blake3"
	"golang.org/x/crypto/argon2"
)

//HashFunc : PoW Hash Function
func HashFunc(data []byte, nonce uint64) []byte {
	h := blake3.New()
	h.Write(data)
	data = h.Sum(nil)
	buf := make([]byte, 8)
	binary.BigEndian.PutUint64(buf, nonce)
	h.Reset()
	h.Write(buf)
	h.Write(data)
	buf = h.Sum(nil)[:16]
	return argon2.IDKey(data, buf, 1, 64*1024, 4, 32)
}
