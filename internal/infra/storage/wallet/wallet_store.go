package storage

type WalletData struct {
	Adress  string
	PubKey  string
	PrivKey string
}

type WalletStore interface {
	SaveWallet(wallet WalletData) error
	LoadWallet(address string) (WalletData, error)
	ListWallets() ([]WalletData, error)
}
