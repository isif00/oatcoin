package oatcoin

import (
	"fmt"

	"github.com/isif00/oat-coin/internal/app"
	"github.com/isif00/oat-coin/internal/infra/filesystem"
	blockstore "github.com/isif00/oat-coin/internal/infra/storage/block"
	walletstore "github.com/isif00/oat-coin/internal/infra/storage/wallet"
)

type OatCoin struct {
	WalletApp *app.WalletApp
	BlockApp  *app.BlockApp
}

func NewOatCoin(dataDir string) (*OatCoin, error) {
	fs, err := filesystem.NewFileSystem(dataDir)
	if err != nil {
		return nil, fmt.Errorf("failed to init filesystem: %w", err)
	}

	walletStore, err := walletstore.NewFileWalletStore(fs)
	if err != nil {
		return nil, fmt.Errorf("failed to init wallet store: %w", err)
	}

	blockStore, err := blockstore.NewFileBlockStore(fs)
	if err != nil {
		return nil, fmt.Errorf("failed to init block store: %w", err)
	}

	return &OatCoin{
		WalletApp: app.NewWalletApp(walletStore),
		BlockApp:  app.NewBlockApp(blockStore),
	}, nil
}
