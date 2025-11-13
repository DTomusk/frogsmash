package repos

type StorageClient struct {
}

func NewStorageClient() *StorageClient {
	return &StorageClient{}
}

func (s *StorageClient) UploadFile() {
	// Implement the logic to upload a file to your storage solution
}
