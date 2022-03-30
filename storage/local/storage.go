// Package local implements basic local file storage logic.
package local

import (
	"context"
	"errors"
	"io"
	"os"
	"path/filepath"

	"github.com/alyakimenko/pageshot/config"
	"github.com/alyakimenko/pageshot/storage"
)

// Storage implement local file storage.
type Storage struct {
	dir string
}

// StorageParams is an incoming params for the NewStorage function.
type StorageParams struct {
	Config config.LocalStorageConfig
}

// NewStorage creates new instance of the Storage.
func NewStorage(params StorageParams) *Storage {
	return &Storage{
		dir: params.Config.Directory,
	}
}

// Upload uploads file to the disk.
func (s *Storage) Upload(ctx context.Context, input storage.UploadInput) error {
	path := filepath.Clean(s.dir + "/" + input.Filename)

	file, err := os.Create(path)
	if err != nil {
		return err
	}

	_, err = io.Copy(file, input.File)
	if err != nil {
		return err
	}

	return nil
}

// Download downloads file from the disk by the provided filename.
func (s *Storage) Download(ctx context.Context, filename string) (io.Reader, error) {
	path := filepath.Clean(s.dir + "/" + filename)

	// check if the file already exists
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		return nil, storage.ErrFileNotExist
	}

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	return file, nil
}
