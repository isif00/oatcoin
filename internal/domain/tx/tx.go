package tx

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"fmt"
)

type TxInput struct {
	TxID      string
	OutputIdx int
	Signature []byte
	PubKey    []byte
}

type TxOutput struct {
	Amount     int
	PubKeyHash []byte
}

type Transaction struct {
	ID      string
	Inputs  []TxInput
	Outputs []TxOutput
}

func NewCoinbaseTx(to string, reward int) (*Transaction, error) {
	output := TxOutput{
		Amount:     reward,
		PubKeyHash: []byte(to),
	}
	tx := &Transaction{
		Inputs:  []TxInput{},
		Outputs: []TxOutput{output},
	}

	id, err := tx.Hash()
	if err != nil {
		return nil, err
	}
	tx.ID = id

	return tx, nil
}

func (tx *Transaction) Hash() (string, error) {
	var buf bytes.Buffer

	enc := gob.NewEncoder(&buf)
	err := enc.Encode(tx)
	if err != nil {
		return "", err
	}

	hash := sha256.Sum256(buf.Bytes())

	return fmt.Sprintf("%x", hash[:]), nil
}
