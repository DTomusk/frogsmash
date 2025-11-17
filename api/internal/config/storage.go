package config

type StorageConfig struct {
	StorageAccountID string
	StorageAccessKey string
	StorageSecretKey string
	StorageBucket    string
}

func NewStorageConfig() *StorageConfig {
	return &StorageConfig{
		StorageAccountID: getEnv("STORAGE_ACCOUNT_ID"),
		StorageAccessKey: getEnv("STORAGE_ACCESS_KEY"),
		StorageSecretKey: getEnv("STORAGE_SECRET_KEY"),
		StorageBucket:    getEnv("STORAGE_BUCKET"),
	}
}
