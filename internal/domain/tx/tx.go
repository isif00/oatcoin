package tx

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"fmt"

	"github.com/decred/dcrd/dcrec/secp256k1/v4"
	"github.com/decred/dcrd/dcrec/secp256k1/v4/ecdsa"
)

type TxInput struct {
	TxID      []byte
	OutputIdx int
	Signature []byte
	PubKey    []byte
}

type TxOutput struct {
	Amount     int
	PubKeyHash []byte
}

type Transaction struct {
	ID      []byte
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

func NewUTXOTransaction(from, to string, amount int, utxos map[string][]TxOutput) (*Transaction, error) {
	var inputs []TxInput
	var outputs []TxOutput

	accumulated := 0
	spentOutputs := make(map[string]int)

	for txid, outs := range utxos {
		for idx, out := range outs {
			if string(out.PubKeyHash) == from {
				accumulated += out.Amount
				inputs = append(inputs, TxInput{
					TxID:      []byte(txid),
					OutputIdx: idx,
					Signature: nil,          // to be signed
					PubKey:    []byte(from), // simplified pubkey
				})
				spentOutputs[txid] = idx
				if accumulated >= amount {
					break
				}
			}
		}
		if accumulated >= amount {
			break
		}
	}

	if accumulated < amount {
		return nil, fmt.Errorf("❌ not enough funds")
	}

	outputs = append(outputs, TxOutput{
		Amount:     amount,
		PubKeyHash: []byte(to),
	})

	if accumulated > amount {
		outputs = append(outputs, TxOutput{
			Amount:     accumulated - amount,
			PubKeyHash: []byte(from),
		})
	}

	tx := &Transaction{Inputs: inputs, Outputs: outputs}
	id, err := tx.Hash()
	if err != nil {
		return nil, err
	}
	tx.ID = id

	return tx, nil
}

func (tx *Transaction) Sign(priv *secp256k1.PrivateKey) error {
	txHash, err := tx.Hash()
	if err != nil {
		return err
	}

	for i := range tx.Inputs {
		sig := ecdsa.Sign(priv, txHash)
		tx.Inputs[i].Signature = sig.Serialize()
		tx.Inputs[i].PubKey = priv.PubKey().SerializeCompressed()
	}
	return nil
}

func (tx *Transaction) Verify() bool {
	txHash, err := tx.Hash()
	if err != nil {
		return false
	}

	for _, in := range tx.Inputs {
		pubKey, err := secp256k1.ParsePubKey(in.PubKey)
		if err != nil {
			return false
		}

		sig, err := ecdsa.ParseDERSignature(in.Signature)
		if err != nil {
			return false
		}

		if !sig.Verify(txHash, pubKey) {
			return false
		}
	}
	return true
}

func (tx *Transaction) Hash() ([]byte, error) {
	var buf bytes.Buffer
	copyTx := *tx
	for i := range copyTx.Inputs {
		copyTx.Inputs[i].Signature = nil
		copyTx.Inputs[i].PubKey = nil
	}
	_ = gob.NewEncoder(&buf).Encode(copyTx)
	hash := sha256.Sum256(buf.Bytes())
	return hash[:], nil
}
