package utxo

import (
	"fmt"
	"github.com/decred/dcrd/dcrec/secp256k1/v4"

	"github.com/isif00/oat-coin/internal/domain/block"
	"github.com/isif00/oat-coin/internal/domain/tx"
)

func CreateUTXOTransaction(fromPriv *secp256k1.PrivateKey, to string, amount int, utxos map[string][]tx.TxOutput) (*tx.Transaction, error) {
	from := string(fromPriv.PubKey().SerializeCompressed())

	tx, err := tx.NewUTXOTransaction(from, to, amount, utxos)
	if err != nil {
		return nil, err
	}

	err = tx.Sign(fromPriv)
	if err != nil {
		return nil, fmt.Errorf("failed to sign transaction: %v", err)
	}

	for _, input := range tx.Inputs {
		txID := string(input.TxID)
		delete(utxos, txID)
	}

	utxos[string(tx.ID)] = tx.Outputs

	return tx, nil
}

func GetUTXOSet(blocks []*block.Block) map[string][]tx.TxOutput {
	utxos := make(map[string][]tx.TxOutput)
	spent := make(map[string]map[int]bool)

	for _, block := range blocks {
		for _, tx := range block.Transactions {
			txID := string(tx.ID)

			// Skip spent outputs
			for idx, out := range tx.Outputs {
				if spent[txID][idx] {
					continue
				}
				utxos[txID] = append(utxos[txID], out)
			}

			// Track spent outputs
			for _, in := range tx.Inputs {
				inTxID := string(in.TxID)
				if spent[inTxID] == nil {
					spent[inTxID] = make(map[int]bool)
				}
				spent[inTxID][in.OutputIdx] = true
			}
		}
	}
	return utxos
}
