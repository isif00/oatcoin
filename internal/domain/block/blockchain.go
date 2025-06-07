package block

import "github.com/isif00/oat-coin/internal/domain/tx"

type Blockchain struct {
	Blocks []*Block
}

func NewBlockchain() *Blockchain {
	genesis := NewBlock([]*tx.Transaction{}, []byte{})
	return &Blockchain{[]*Block{genesis}}
}

func (bc *Blockchain) AddBlock(txs []*tx.Transaction) {
	prev := bc.Blocks[len(bc.Blocks)-1]
	newBlock := NewBlock(txs, prev.Hash)
	bc.Blocks = append(bc.Blocks, newBlock)
}
