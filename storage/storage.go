package storage

import "io"

// UploadInput is an input model for uploading files.
type UploadInput struct {
	Filename string
	File     io.Reader
}
