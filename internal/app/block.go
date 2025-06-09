package app

import (
	"github.com/isif00/oat-coin/internal/domain/block"
	"github.com/isif00/oat-coin/internal/domain/tx"
	storage "github.com/isif00/oat-coin/internal/infra/storage/block"
)

type BlockApp struct {
	store storage.BlockStore
}

func NewBlockApp(store storage.BlockStore) *BlockApp {
	return &BlockApp{store: store}
}

func InitializeBlockchain(store storage.BlockStore) ([]*block.Block, error) {
	return block.InitializeBlockchain(store)
}

func MineBlock(store storage.BlockStore, txs []*tx.Transaction) (*block.Block, error) {
	return block.MineBlock(store, txs)
}

func GetLatestBlock(store storage.BlockStore) (*block.Block, error) {
	return block.GetLatestBlock(store)
}

func GetAllBlocks(store storage.BlockStore) ([]*block.Block, error) {
	return block.GetAllBlocks(store)
}
