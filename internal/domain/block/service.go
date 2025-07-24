package block

import (
	"github.com/isif00/oat-coin/internal/domain/tx"
	storage "github.com/isif00/oat-coin/internal/infra/storage/block"
)

func InitializeBlockchain(store storage.BlockStore) ([]*Block, error) {
	blocks, err := store.LoadAllBlocks()
	if err != nil {
		return nil, err
	}

	domainBlocks := make([]*Block, 0, len(blocks))
	for _, bd := range blocks {
		domainBlocks = append(domainBlocks, ToDomain(bd))
	}

	if len(domainBlocks) == 0 {
		genesis := NewBlock([]*tx.Transaction{}, []byte{})
		genesisData := genesis.ToStorage()
		if err := store.SaveBlock(genesisData); err != nil {
			return nil, err
		}
		domainBlocks = append(domainBlocks, genesis)
	}

	return domainBlocks, nil
}

func MineBlock(store storage.BlockStore, txs []*tx.Transaction) (*Block, error) {
	currentBlocks, err := store.LoadAllBlocks()
	if err != nil {
		return nil, err
	}

	var prevHash []byte
	if len(currentBlocks) > 0 {
		prevHash = []byte(currentBlocks[len(currentBlocks)-1].Hash)
	}

	newBlock := NewBlock(txs, prevHash)
	storageBlock := newBlock.ToStorage()

	if err := store.SaveBlock(storageBlock); err != nil {
		return nil, err
	}
	return newBlock, nil
}

func GetLatestBlock(store storage.BlockStore) (*Block, error) {
	blocks, err := store.LoadAllBlocks()
	if err != nil {
		return nil, err
	}

	if len(blocks) == 0 {
		return nil, nil
	}

	latest := blocks[len(blocks)-1]
	return ToDomain(latest), nil
}

func GetAllBlocks(store storage.BlockStore) ([]*Block, error) {
	rawBlocks, err := store.LoadAllBlocks()
	if err != nil {
		return nil, err
	}

	domainBlocks := make([]*Block, 0, len(rawBlocks))
	for _, bd := range rawBlocks {
		domainBlocks = append(domainBlocks, ToDomain(bd))
	}

	return domainBlocks, nil
}
