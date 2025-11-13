package repos

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type StorageClient struct {
	S3Client *s3.Client
	Bucket   string
	Endpoint string
}

func NewStorageClient(ctx context.Context, accountID, accessKey, secretKey, bucket string) (*StorageClient, error) {
	endpoint := "https://" + accountID + ".r2.cloudflarestorage.com"

	cfg, err := config.LoadDefaultConfig(
		ctx, config.WithCredentialsProvider(
			credentials.NewStaticCredentialsProvider(accessKey, secretKey, ""),
		),
	)

	if err != nil {
		return nil, err
	}

	client := s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.BaseEndpoint = aws.String(endpoint)
	})

	return &StorageClient{
		S3Client: client,
		Bucket:   bucket,
		Endpoint: endpoint,
	}, nil
}

func (s *StorageClient) UploadFile() {
	// Implement the logic to upload a file to your storage solution
}
