package block

import (
	"bytes"
	"encoding/gob"
	"runtime"
	"time"

	"github.com/isif00/oat-coin/internal/domain/tx"
)

type Block struct {
	Timestamp    int64
	Transactions []*tx.Transaction
	PrevHash     []byte
	Hash         []byte
	Nonce        int
}

func (b *Block) Serialize() []byte {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	_ = enc.Encode(b)
	return buf.Bytes()
}

func Deserialize(data []byte) *Block {
	var block Block
	dec := gob.NewDecoder(bytes.NewReader(data))
	_ = dec.Decode(&block)
	return &block
}

func NewBlock(txs []*tx.Transaction, prevHash []byte) *Block {
	block := &Block{
		Timestamp:    time.Now().Unix(),
		Transactions: txs,
		PrevHash:     prevHash,
		Hash:         []byte{},
	}
	pow := NewProofOfWork(block)
	nonce, hash := pow.Run(runtime.NumCPU())
	block.Hash = hash[:]
	block.Nonce = nonce
	return block
}
