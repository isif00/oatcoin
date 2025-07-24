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

func (b *BlockApp) InitializeBlockchain() ([]*block.Block, error) {
	return block.InitializeBlockchain(b.store)
}

func (b *BlockApp) MineBlock(txs []*tx.Transaction) (*block.Block, error) {
	return block.MineBlock(b.store, txs)
}

func (b *BlockApp) GetLatestBlock() (*block.Block, error) {
	return block.GetLatestBlock(b.store)
}

func (b *BlockApp) GetAllBlocks() ([]*block.Block, error) {
	return block.GetAllBlocks(b.store)
}
