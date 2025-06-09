package block

import (
	"encoding/hex"
	"runtime"
	"time"

	"github.com/isif00/oat-coin/internal/domain/tx"
	"github.com/isif00/oat-coin/internal/infra/storage/block"
)

type Block struct {
	Nonce        int
	Hash         []byte
	PrevHash     []byte
	Timestamp    int64
	Transactions []*tx.Transaction
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

func (b *Block) ToStorage() block.BlockData {
	txIDs := make([]string, len(b.Transactions))
	for i, tx := range b.Transactions {
		txIDs[i] = hex.EncodeToString([]byte(tx.ID))
	}
	return block.BlockData{
		Hash:         hex.EncodeToString(b.Hash),
		PrevHash:     hex.EncodeToString(b.PrevHash),
		Timestamp:    b.Timestamp,
		Nonce:        b.Nonce,
		Transactions: txIDs,
	}
}

func BlockDataToDomain(bd block.BlockData) *Block {
	txHashes := make([]*tx.Transaction, len(bd.Transactions))
	for i, id := range bd.Transactions {
		id, _ := hex.DecodeString(id)
		txHashes[i] = &tx.Transaction{ID: hex.EncodeToString(id)}
	}
	return &Block{
		Hash:         []byte(bd.Hash),
		PrevHash:     []byte(bd.PrevHash),
		Timestamp:    bd.Timestamp,
		Nonce:        bd.Nonce,
		Transactions: txHashes,
	}
}

func ToDomain(bd block.BlockData) *Block {
	txList := make([]*tx.Transaction, len(bd.Transactions))
	for i, id := range bd.Transactions {
		txList[i] = &tx.Transaction{
			ID: id,
		}
	}

	return &Block{
		Hash:         []byte(bd.Hash),
		PrevHash:     []byte(bd.PrevHash),
		Timestamp:    bd.Timestamp,
		Nonce:        bd.Nonce,
		Transactions: txList,
	}
}
