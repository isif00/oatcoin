package filesystem

import (
	"os"
	"path/filepath"
	"strings"
)

type FileSystem struct {
	BasePath string
}

func NewFileSystem(root string) (*FileSystem, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	fullPath := filepath.Join(home, root)
	if err := os.MkdirAll(fullPath, 0700); err != nil {
		return nil, err
	}
	return &FileSystem{BasePath: fullPath}, nil
}

func (fs *FileSystem) Write(folder, file string, data []byte) error {
	path := filepath.Join(fs.BasePath, folder)
	if err := os.MkdirAll(path, 0700); err != nil {
		return err
	}
	return os.WriteFile(filepath.Join(path, file), data, 0600)
}

func (fs *FileSystem) Read(folder, file string) ([]byte, error) {
	return os.ReadFile(filepath.Join(fs.BasePath, folder, file))
}

func (fs *FileSystem) ListDirs() ([]string, error) {
	files, err := os.ReadDir(fs.BasePath)
	if err != nil {
		return nil, err
	}
	var dirs []string
	for _, f := range files {
		if f.IsDir() {
			dirs = append(dirs, f.Name())
		}
	}
	return dirs, nil
}

func (f *FileSystem) ListFiles(folder string) ([]string, error) {
	fullPath := filepath.Join(f.BasePath, folder)
	entries, err := os.ReadDir(fullPath)
	if err != nil {
		return nil, err
	}
	var files []string
	for _, e := range entries {
		if !e.IsDir() && strings.HasSuffix(e.Name(), ".json") {
			files = append(files, strings.TrimSuffix(e.Name(), ".json"))
		}
	}
	return files, nil
}
