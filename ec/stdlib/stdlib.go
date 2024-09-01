package stdlib

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/md5"
	"crypto/rand"
	"io"
	"math/big"
)

func generateKey() (*ecdsa.PrivateKey, *ecdsa.PublicKey, error) {
	privkey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return nil, nil, err
	}
	pubkey := &privkey.PublicKey

	return privkey, pubkey, nil
}

// The signature are in the form of big ints r & s
func signMessage(msg string, privkey *ecdsa.PrivateKey) (r *big.Int, s *big.Int, signhash []byte, err error) {
	h := md5.New()
	io.WriteString(h, msg)
	signhash = h.Sum(nil) // convert hash.Hash to byte slice

	r, s, err = ecdsa.Sign(rand.Reader, privkey, signhash)
	return
}

func verifyMessage(r *big.Int, s *big.Int, pubkey *ecdsa.PublicKey, signhash []byte) bool {
	return ecdsa.Verify(pubkey, signhash, r, s)
}
