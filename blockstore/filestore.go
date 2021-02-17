package blockstore

import (
	"encoding/hex"
	"os"
	"path/filepath"

	"github.com/zeebo/blake3"
)

//FileStore : Save block to file
type FileStore struct {
	root string
}

//InitFileStore :
func InitFileStore(root string) *FileStore {
	for i := byte(0); true; i++ {
		os.MkdirAll(filepath.Join(root, hex.EncodeToString([]byte{i})), os.ModeDir)
		if i == 255 {
			break
		}
	}
	f := new(FileStore)
	f.root = root
	return f
}

//Store : create new store object
func (fs *FileStore) Store(filename string) (*os.File, error) {
	h := blake3.New()
	h.WriteString(filename)
	filename = hex.EncodeToString(h.Sum(nil))
	return os.OpenFile(filepath.Join(fs.root, filename[:2], filename), os.O_CREATE, os.ModeAppend)
}

//Get : create new store object
func (fs *FileStore) Get(filename string) (*os.File, error) {
	h := blake3.New()
	h.WriteString(filename)
	filename = hex.EncodeToString(h.Sum(nil))
	return os.Open(filepath.Join(fs.root, filename[:2], filename))
}
