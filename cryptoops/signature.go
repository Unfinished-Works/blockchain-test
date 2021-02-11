package cryptoops

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"errors"
	"math/big"
	"strings"
)

//ErrPublicKeyInvaild : This error occurs when the given public key is invalid.
var ErrPublicKeyInvaild error = errors.New("PublicKey Invaild")
var p384 elliptic.Curve = elliptic.P384()

//LoadPublicKey : Load the public key string.
func LoadPublicKey(key string) (*ecdsa.PublicKey, error) {
	parts := strings.Split(key, "$")
	var x, y big.Int
	_, ok := x.SetString(parts[0], 62)
	if !ok {
		return nil, ErrPublicKeyInvaild
	}
	_, ok = y.SetString(parts[1], 62)
	if !ok {
		return nil, ErrPublicKeyInvaild
	}
	if !p384.IsOnCurve(&x, &y) {
		return nil, ErrPublicKeyInvaild
	}
	publickey := new(ecdsa.PublicKey)
	publickey.Curve = p384
	publickey.X = &x
	publickey.Y = &y
	return publickey, nil
}
