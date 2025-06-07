package storage

import (
	"encoding/json"

	"github.com/isif00/oat-coin/internal/infra/filesystem"
)

type FileWalletStore struct {
	fs     *filesystem.FileSystem
	folder string
}

func NewFileWalletStore(fs *filesystem.FileSystem) *FileWalletStore {
	return &FileWalletStore{
		fs:     fs,
		folder: "wallets",
	}
}

func (s *FileWalletStore) SaveWallet(wallet WalletData) error {
	data, err := json.Marshal(wallet)
	if err != nil {
		return err
	}
	return s.fs.Write(s.folder, wallet.Adress+".json", data)
}

func (s *FileWalletStore) LoadWallet(address string) (WalletData, error) {
	data, err := s.fs.Read(s.folder, address+".json")
	if err != nil {
		return WalletData{}, err
	}
	var wallet WalletData
	if err := json.Unmarshal(data, &wallet); err != nil {
		return WalletData{}, err
	}
	return wallet, nil
}

func (s *FileWalletStore) ListWallets() ([]WalletData, error) {
	files, err := s.fs.ListFiles(s.folder)
	if err != nil {
		return nil, err
	}

	var wallets []WalletData
	for _, addr := range files {
		wallet, err := s.LoadWallet(addr)
		if err == nil {
			wallets = append(wallets, wallet)
		}
	}	
	return wallets, nil
}
