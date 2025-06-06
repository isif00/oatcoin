package wallet

import (
	"github.com/decred/dcrd/dcrec/secp256k1/v4"
	"github.com/isif00/oat-coin/pkg/crypto"
)

type Wallet struct {
	PrivateKey secp256k1.PrivateKey
	PublicKey  secp256k1.PublicKey
}

func NewWallet() (*Wallet, error) {
	privKey, err := crypto.GenerateKey()
	if err != nil {
		return nil, err
	}

	pubKey := privKey.PubKey()

	return &Wallet{
		PrivateKey: *privKey,
		PublicKey:  *pubKey,
	}, nil
}

func (w *Wallet) Address() string {
	pubKeyHash := crypto.HashPubKey(w.PublicKey.SerializeCompressed())

	versionedPayload := append([]byte{0x00}, pubKeyHash...)

	address := crypto.Base58CheckEncode(versionedPayload)

	return address

}
