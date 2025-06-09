package oatcoin

import (
	"log"

	"github.com/isif00/oat-coin/internal/app"
	"github.com/isif00/oat-coin/internal/infra/filesystem"
	blockstore "github.com/isif00/oat-coin/internal/infra/storage/block"
	walletstore "github.com/isif00/oat-coin/internal/infra/storage/wallet"
)

var (
	fs *filesystem.FileSystem

	walletStore *walletstore.FileWalletStore
	blockStore  *blockstore.FileBlockStore

	walletApp *app.WalletApp
	blockApp  *app.BlockApp
)

func init() {
	var err error
	fs, err = filesystem.NewFileSystem(".oatcoin")
	if err != nil {
		log.Fatalf("failed to init filesystem: %v", err)
	}

	walletStore = walletstore.NewFileWalletStore(fs)
	walletApp = app.NewWalletApp(walletStore)

	blockStore = blockstore.NewFileBlockStore(fs)
	blockApp = app.NewBlockApp(blockStore)
}
