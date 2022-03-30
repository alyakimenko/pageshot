// Package s3 implements S3 storage logic.
package s3

import (
	"context"
	"io"

	"github.com/alyakimenko/pageshot/config"
	"github.com/alyakimenko/pageshot/storage"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

// Storage is a S3 implementation of a file storage.
type Storage struct {
	client *minio.Client
	bucket string
}

// StorageParams is an incoming params for the NewStorage function.
type StorageParams struct {
	Config config.S3StorageConfig
}

// NewStorage creates new instance of the Storage.
func NewStorage(params StorageParams) (*Storage, error) {
	client, err := minio.New(params.Config.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(params.Config.AccessKeyID, params.Config.SecretAccessKey, ""),
		Secure: params.Config.SSL,
	})
	if err != nil {
		return nil, err
	}

	return &Storage{
		client: client,
		bucket: params.Config.Bucket,
	}, nil
}

// Upload uploads file to S3.
func (s *Storage) Upload(ctx context.Context, input storage.UploadInput) error {
	opts := minio.PutObjectOptions{
		ContentType:  input.ContentType,
		UserMetadata: map[string]string{"x-amz-acl": "public-read"},
	}

	_, err := s.client.PutObject(ctx, s.bucket, input.Filename, input.File, input.FileSize, opts)
	if err != nil {
		return err
	}

	return nil
}

// Download downloads file from S3 by the provided filename.
func (s *Storage) Download(ctx context.Context, filename string) (io.Reader, error) {
	object, err := s.client.GetObject(ctx, s.bucket, filename, minio.GetObjectOptions{})
	if err != nil {
		return nil, err
	}

	return object, nil
}
