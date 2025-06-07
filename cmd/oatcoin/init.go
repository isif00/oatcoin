package oatcoin

import (
	"log"

	"github.com/isif00/oat-coin/internal/app"
	"github.com/isif00/oat-coin/internal/infra/filesystem"
	storage "github.com/isif00/oat-coin/internal/infra/storage/wallet"
)

var (
	fs        *filesystem.FileSystem
	store     *storage.FileWalletStore
	walletApp *app.WalletApp
)

func init() {
	var err error
	fs, err = filesystem.NewFileSystem(".oatcoin")
	if err != nil {
		log.Fatalf("failed to init filesystem: %v", err)
	}

	// wallet
	store = storage.NewFileWalletStore(fs)
	walletApp = app.NewWalletApp(store)
}
