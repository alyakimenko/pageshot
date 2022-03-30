package config

// StorageType represents a storage type.
type StorageType string

const (
	// LocalStorageType represents local storage type.
	LocalStorageType StorageType = "local"
	// S3StorageType represents S3 storage type.
	S3StorageType StorageType = "s3"
)

// IsValid validates the StorageType.
func (st StorageType) IsValid() bool {
	return st == LocalStorageType || st == S3StorageType
}

// StorageConfig holds storage related configurable values.
type StorageConfig struct {
	Type  StorageType
	Local LocalStorageConfig
	S3    S3StorageConfig
}

// LocalStorageConfig holds local storage related configurable values.
type LocalStorageConfig struct {
	Directory string
}

// S3StorageConfig holds S3 storage related configurable values.
type S3StorageConfig struct {
	Bucket          string
	Endpoint        string
	AccessKeyID     string
	SecretAccessKey string
	SSL             bool
}
