package wallet

import (
	"encoding/hex"

	storage "github.com/isif00/oat-coin/internal/infra/storage/wallet"
)

func CreateWallet(store storage.WalletStore) (string, error) {
	w, err := NewWallet()
	a := w.Address()
	if err != nil {
		return "", err
	}

	data := storage.WalletData{
		Adress:  a,
		PubKey:  hex.EncodeToString(w.PublicKey.SerializeCompressed()),
		PrivKey: hex.EncodeToString(w.PrivateKey.Serialize()),
	}

	if err := store.SaveWallet(data); err != nil {
		return "", err
	}
	return a, nil
}

func LoadWallet(address string, store storage.WalletStore) (storage.WalletData, error) {
	return store.LoadWallet(address)

}

func ListWallets(store storage.WalletStore) ([]storage.WalletData, error) {
	return store.ListWallets()
}
