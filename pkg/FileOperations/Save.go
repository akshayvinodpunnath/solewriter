package FileOperations

import (
	"os"
	"path/filepath"
)

type FileOperation struct {
	Input    []byte
	FileName string
	Path     string
}

func (fo *FileOperation) SaveFile() error {
	fullPath := filepath.Join(fo.Path, fo.FileName)

	err := os.MkdirAll(fo.Path, 0755)
	if err != nil {
		return err
	}

	err = os.WriteFile(fullPath, fo.Input, 0644)
	if err != nil {
		return err
	}
	return nil
}

func (fo *FileOperation) ReadFile() ([]byte, error) {
	fullPath := filepath.Join(fo.Path, fo.FileName)

	data, err := os.ReadFile(fullPath)
	if err != nil {
		return nil, err
	}
	return data, nil
}
