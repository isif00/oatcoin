package app

import (
	"github.com/isif00/oat-coin/internal/domain/wallet"
	storage "github.com/isif00/oat-coin/internal/infra/storage/wallet"
)

type WalletApp struct {
	store storage.WalletStore
}

func NewWalletApp(store storage.WalletStore) *WalletApp {
	return &WalletApp{store: store}
}

func (wa *WalletApp) CreateWallet() (string, error) {
	return wallet.CreateWallet(wa.store)
}

func (wa *WalletApp) LoadWallet(address string) (storage.WalletData, error) {
	return wallet.LoadWallet(address, wa.store)
}

func (wa *WalletApp) ListWallets() ([]storage.WalletData, error) {
	return wallet.ListWallets(wa.store)
}
