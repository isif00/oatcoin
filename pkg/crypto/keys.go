package crypto

import "github.com/decred/dcrd/dcrec/secp256k1/v4"

func GenerateKey() (*secp256k1.PrivateKey, error) {
	return secp256k1.GeneratePrivateKey()
}
