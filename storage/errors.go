package storage

import "errors"

var (
	// ErrFileNotExist occurs when trying to access a non-existent file.
	ErrFileNotExist = errors.New("file does not exist")
)
