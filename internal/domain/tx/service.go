package tx

import "github.com/decred/dcrd/dcrec/secp256k1/v4"

func CreateCoinbaseTransaction(to string, reward int) (*Transaction, error) {
	return NewCoinbaseTx(to, reward)
}

func SignTransaction(tx *Transaction, priv *secp256k1.PrivateKey) error {
	return tx.Sign(priv)
}
