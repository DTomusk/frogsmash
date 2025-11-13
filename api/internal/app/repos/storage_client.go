package repos

import (
	"context"
	"mime/multipart"

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
		ctx,
		config.WithRegion("auto"),
		config.WithCredentialsProvider(
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

// TODO: should this specifically be a gin context?
func (s *StorageClient) UploadFile(filename string, fileHeader *multipart.FileHeader, ctx context.Context) (string, error) {
	file, err := fileHeader.Open()
	if err != nil {
		return "", err
	}
	defer file.Close()

	_, err = s.S3Client.PutObject(ctx, &s3.PutObjectInput{
		Bucket: aws.String(s.Bucket),
		Key:    aws.String(filename),
		Body:   file,
	})

	if err != nil {
		return "", err
	}

	return s.Endpoint + "/" + filename, nil
}
