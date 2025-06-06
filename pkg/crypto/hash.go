package crypto

import (
	"crypto/sha256"

	"golang.org/x/crypto/ripemd160"
)

func HashPubKey(pubKey []byte) []byte {
	shaHash := sha256.Sum256(pubKey)

	ripemd := ripemd160.New()
	ripemd.Write(shaHash[:])

	return ripemd.Sum(nil)
}
