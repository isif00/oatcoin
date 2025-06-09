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

func NewFileBlockStore(fs *filesystem.FileSystem) (*FileBlockStore, error) {
	store := &FileBlockStore{
		fs:     fs,
		folder: "blocks",
	}

	dirPath := filepath.Join(fs.BasePath, "blocks")
	if err := os.MkdirAll(dirPath, os.ModePerm); err != nil {
		return nil, err
	}

	return store, nil
}

func (s *FileBlockStore) SaveBlock(b BlockData) error {
	data, err := json.Marshal(b)
	if err != nil {
		return err
	}
	filename := b.Hash + ".json"
	err = s.fs.Write(s.folder, filename, data)
	if err != nil {
		return err
	}
	return err
}

func (s *FileBlockStore) LoadBlock(hash string) (BlockData, error) {
	data, err := s.fs.Read(s.folder, hash+".json")
	if err != nil {
		return BlockData{}, err
	}

	var bd BlockData
	err = json.Unmarshal(data, &bd)
	if err != nil {
		return BlockData{}, err
	}
	return bd, err
}

func (s *FileBlockStore) LoadAllBlocks() ([]BlockData, error) {
	names, err := s.fs.ListFiles(s.folder)
	if err != nil {
		return nil, err
	}
	var blocks []BlockData
	for _, name := range names {
		blk, err := s.LoadBlock(name)
		if err != nil {
			return nil, err
		}
		blocks = append(blocks, blk)
	}
	return blocks, nil
}
