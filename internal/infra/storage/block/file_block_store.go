package block

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/isif00/oat-coin/internal/infra/filesystem"
)

type FileBlockStore struct {
	fs     *filesystem.FileSystem
	folder string
}

func NewFileBlockStore(fs *filesystem.FileSystem) *FileBlockStore {
	store := &FileBlockStore{
		fs:     fs,
		folder: "blocks",
	}

	_ = os.MkdirAll(filepath.Join(fs.BasePath, "blocks"), os.ModePerm)

	return store
}

func (s *FileBlockStore) SaveBlock(b BlockData) error {
	data, err := json.Marshal(b)
	if err != nil {
		return err
	}
	filename := b.Hash + ".json"
	return s.fs.Write(s.folder, filename, data)
}

func (s *FileBlockStore) LoadBlock(hash string) (BlockData, error) {
	data, err := s.fs.Read(s.folder, hash+".json")
	if err != nil {
		return BlockData{}, err
	}
	var bd BlockData
	err = json.Unmarshal(data, &bd)
	return bd, err
}

func (s *FileBlockStore) LoadAllBlocks() ([]BlockData, error) {
	names, err := s.fs.ListFiles(s.folder)
	if err != nil {
		return nil, err
	}
	var blocks []BlockData
	for _, name := range names {
		blk, err := s.LoadBlock(name[:len(name)-5])
		if err != nil {
			continue
		}
		blocks = append(blocks, blk)
	}
	return blocks, nil
}
