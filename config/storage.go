package config

// StorageType represents a storage type.
type StorageType string

const (
	// LocalStorageType represents local storage type.
	LocalStorageType StorageType = "local"
)

// IsValid validates the StorageType.
func (st StorageType) IsValid() bool {
	return st == LocalStorageType
}

// StorageConfig holds storage related configurable values.
type StorageConfig struct {
	Type      StorageType
	Directory string
}
