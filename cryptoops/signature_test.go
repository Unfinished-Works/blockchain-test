package cryptoops

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"reflect"
	"testing"
)

func TestPublicKeySaveLoad(t *testing.T) {
	privkey, _ := ecdsa.GenerateKey(elliptic.P384(), rand.Reader)
	publickey := privkey.PublicKey
	pubstring := ExportPublicKey(&publickey)
	//fmt.Println(pubstring)
	t.Run("", func(t *testing.T) {
		got, err := LoadPublicKey(pubstring)
		if err != nil {
			t.Errorf("LoadPublicKey() error = %v, wantErr %v", err, false)
			return
		}
		if !reflect.DeepEqual(*got, publickey) {
			t.Errorf("LoadPublicKey() = %v, want %v", *got, publickey)
		}
	})
}
