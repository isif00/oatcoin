package block

type BlockData struct {
	Nonce        int
	Hash         string
	PrevHash     string
	Timestamp    int64
	Transactions []string
}
type BlockStore interface {
	SaveBlock(b BlockData) error
	LoadBlock(hash string) (BlockData, error)
	LoadAllBlocks() ([]BlockData, error)
}
