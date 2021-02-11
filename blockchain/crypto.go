package blockchain

import (
	"encoding/binary"
	"errors"

	"github.com/zeebo/blake3"
)

func (x *Transaction) verifyStructure() bool {
	var neededFee uint64 = 16
	if len(x.GetRecipient()) < 64 {
		return false
	}
	neededFee += uint64(len(x.GetPayload()) * 16)
	if len(x.GetPayload()) > 8192 {
		return false
	}
	if neededFee > x.GetFee() {
		return false
	}
	return true
}

//ErrInvaildTransaction : Invaild Transaction
var ErrInvaildTransaction = errors.New("Invaild Transaction")

//Hash : Compute the blake3 hash.
func (x *Transaction) Hash() ([]byte, error) {
	if !x.verifyStructure() {
		return nil, ErrInvaildTransaction
	}
	buf := make([]byte, 24)
	h := blake3.New()
	binary.LittleEndian.PutUint64(buf[0:8], x.GetAccountNonce())
	binary.LittleEndian.PutUint64(buf[8:16], x.GetAmount())
	binary.LittleEndian.PutUint64(buf[16:24], x.GetFee())
	h.Write(buf)
	h.Write(x.GetPayload())
	h.Write([]byte(x.GetRecipient()))
	return h.Sum(nil), nil
}
